// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAttributionInventory(t *testing.T) {
	source := filepath.Join("..", "..")
	if _, err := verify(source); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name, path, content, message string
		remove                       bool
	}{
		{name: "missing license", path: "LICENSE", message: "missing inventory", remove: true},
		{name: "altered license", path: "uap-core/LICENSE", content: "wrong owner", message: "original attribution/source changed"},
		{name: "altered copyright", path: "uaparser/LICENSE.md", content: "wrong owner", message: "original attribution/source changed"},
		{name: "altered original", path: "uaparser/cache.go", content: "package uaparser", message: "original attribution/source changed"},
		{name: "unreviewed maintained source", path: "uaparser/parser.go", content: "package uaparser", message: "original attribution/source changed"},
		{name: "altered inventory", path: "UPSTREAM-SHA256SUMS", content: "", message: "original manifest changed"},
		{name: "unaccounted source", path: "new.go", content: "package main", message: "unaccounted file"},
		{name: "unaccounted asset", path: "new.txt", content: "unregistered", message: "unaccounted file"},
		{name: "owned omission", path: "OWNED-FILES", content: "OWNED-FILES\n", message: "unaccounted file"},
		{name: "upstream claimed as owned", path: "OWNED-FILES", content: "LICENSE\n", message: "overlapping owned"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			err := filepath.Walk(source, func(path string, info os.FileInfo, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}
				rel, err := filepath.Rel(source, path)
				if err != nil {
					return err
				}
				target := filepath.Join(root, rel)
				if info.IsDir() {
					return os.MkdirAll(target, 0o755)
				}
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				return ioutil.WriteFile(target, data, 0o644)
			})
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(root, tc.path)
			if tc.remove {
				err = os.Remove(path)
			} else {
				err = ioutil.WriteFile(path, []byte(tc.content), 0o644)
			}
			if err != nil {
				t.Fatal(err)
			}
			if _, err = verify(root); err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("expected %q, got %v", tc.message, err)
			}
		})
	}
}
