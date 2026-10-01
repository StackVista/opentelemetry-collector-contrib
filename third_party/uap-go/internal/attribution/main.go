// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Command attribution verifies copied source before the native first-party
// header checker receives the independent inventory of owned source files.
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// These manifests describe the independently verified original distributions.
var manifests = map[string]string{
	"UPSTREAM-SHA256SUMS": "4edbb43eb34d968093df5bb5134776616c85dcc957b6989bc318e62152be7ba6",
	"UAP-CORE-SHA256SUMS": "bda6f80276a47e85206b61d1e47767bd0c12727a2e8c2d819902c165aa14d60f",
}

// Only the reviewed parser/metadata and cosmetic maintenance deltas are allowed.
// All other original bytes, including copyright and licenses, remain exact.
var maintained = map[string]string{
	"go.mod":                     "49b8b6aa858b829c97726df003eb6fbcc94d90b2f8a81f42fa9234fd0d8abd70",
	"go.sum":                     "bceacfc9b2728035e4098d7c22c311e8555eb923552209f1e329b889f483b643",
	"test.go":                    "91f41949454d562227b4d6be429ffd910630cef283c8dee4cb2655a9b9638a9c",
	"uaparser/benchmark_test.go": "fcb826e49c83ee8abc42bff382f76c87f41705596f50ec5bdb40db6001a36b35",
	"uaparser/parser.go":         "7583d37d7213c6e1e38cecf341d09789e3a70b56c68e08710ea695f3a30a8d85",
	"uaparser/parsing_test.go":   "ed8247e0a98852650681aa722fa19e0c71ed69f5fac148e665b2a16c874d266a",
}

func main() {
	ownedSource := flag.Bool("owned-source", false, "print owned Go/shell paths for the existing header checker")
	flag.Parse()
	owned, err := verify(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *ownedSource {
		for _, name := range owned {
			if strings.HasSuffix(name, ".go") || strings.HasSuffix(name, ".sh") {
				fmt.Println("./" + name)
			}
		}
	} else {
		fmt.Println("Third-party attribution and owned inventory verified")
	}
}

func safePath(name string) bool {
	return name != "." && filepath.ToSlash(filepath.Clean(name)) == name &&
		!filepath.IsAbs(name) && name != ".." && !strings.HasPrefix(name, "../") &&
		!strings.ContainsAny(name, "\\ \t\r\n:")
}

func digest(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func verify(root string) ([]string, error) {
	original := make(map[string]string)
	for manifest, expected := range manifests {
		data, err := ioutil.ReadFile(filepath.Join(root, manifest))
		if err != nil {
			return nil, err
		}
		if digest(data) != expected {
			return nil, fmt.Errorf("original manifest changed: %s", manifest)
		}
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			if len(fields) != 2 || len(fields[0]) != 64 || !safePath(fields[1]) {
				return nil, fmt.Errorf("invalid manifest: %s", manifest)
			}
			if _, exists := original[fields[1]]; exists {
				return nil, fmt.Errorf("duplicate original: %s", fields[1])
			}
			original[fields[1]] = fields[0]
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	for name, expected := range maintained {
		if _, exists := original[name]; !exists {
			return nil, fmt.Errorf("maintenance path is not original: %s", name)
		}
		original[name] = expected
	}
	data, err := ioutil.ReadFile(filepath.Join(root, "OWNED-FILES"))
	if err != nil {
		return nil, err
	}
	accounted := make(map[string]bool)
	for name := range original {
		accounted[name] = true
	}
	var owned []string
	for _, name := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if !safePath(name) || accounted[name] {
			return nil, fmt.Errorf("invalid or overlapping owned inventory: %s", name)
		}
		accounted[name] = true
		owned = append(owned, name)
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		name, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return relErr
		}
		name = filepath.ToSlash(name)
		if !accounted[name] {
			return fmt.Errorf("unaccounted file: %s", name)
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular inventory file: %s", name)
		}
		delete(accounted, name)
		if expected, upstream := original[name]; upstream {
			content, readErr := ioutil.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			if digest(content) != expected {
				return fmt.Errorf("original attribution/source changed: %s", name)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(accounted) != 0 {
		var missing []string
		for name := range accounted {
			missing = append(missing, name)
		}
		sort.Strings(missing)
		return nil, fmt.Errorf("missing inventory files: %s", strings.Join(missing, ", "))
	}
	sort.Strings(owned)
	return owned, nil
}
