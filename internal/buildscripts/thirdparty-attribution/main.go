// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Command attribution verifies copied source before the native first-party
// header checker receives the independent inventory of owned source files.
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type policy struct {
	Manifests  map[string]string
	Maintained map[string]string
}

func main() {
	ownedSource := flag.Bool("owned-source", false, "print owned Go/shell paths for the existing header checker")
	ownedDocs := flag.Bool("owned-docs", false, "print owned source/document paths for the existing spell checker")
	policySHA := flag.String("policy-sha", "", "required reviewed policy SHA-256")
	flag.Parse()
	data, err := ioutil.ReadFile("ATTRIBUTION-POLICY.json")
	if err != nil || len(*policySHA) != 64 || digest(data) != *policySHA {
		fmt.Fprintln(os.Stderr, "missing or changed reviewed attribution policy", err)
		os.Exit(1)
	}
	var p policy
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&p); err != nil || len(p.Manifests) == 0 {
		fmt.Fprintln(os.Stderr, "invalid attribution policy", err)
		os.Exit(1)
	}
	owned, err := verify(".", p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *ownedSource || *ownedDocs {
		for _, name := range owned {
			if strings.HasSuffix(name, ".go") || (*ownedSource && strings.HasSuffix(name, ".sh")) ||
				(*ownedDocs && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".yaml"))) {
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

func verify(root string, p policy) ([]string, error) {
	original := make(map[string]string)
	for manifest, expected := range p.Manifests {
		data, err := ioutil.ReadFile(filepath.Join(root, manifest))
		if err != nil {
			return nil, err
		}
		if digest(data) != expected {
			return nil, fmt.Errorf("original manifest changed: %s", manifest)
		}
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		for scanner.Scan() {
			fields := strings.SplitN(scanner.Text(), "  ", 2)
			if len(fields) != 2 || len(fields[0]) != 64 || fields[1] == "." || filepath.ToSlash(filepath.Clean(fields[1])) != fields[1] || filepath.IsAbs(fields[1]) || strings.HasPrefix(fields[1], "../") {
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
	for name, expected := range p.Maintained {
		if name != "go.mod" && name != "go.sum" && !strings.HasSuffix(name, ".go") {
			return nil, fmt.Errorf("maintenance exception is not Go source/metadata: %s", name)
		}
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
