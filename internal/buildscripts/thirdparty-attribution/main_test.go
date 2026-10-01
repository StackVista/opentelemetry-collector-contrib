// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOriginalAndOwnedBoundaries(t *testing.T) {
	cases := []struct {
		name, path, content, message string
		remove, licenseException     bool
	}{
		{name: "valid"},
		{name: "missing license", path: "LICENSE", remove: true, message: "missing inventory"},
		{name: "altered license", path: "LICENSE", content: "new owner", message: "original attribution/source changed"},
		{name: "altered copyright", path: "COPYRIGHT", content: "new owner", message: "original attribution/source changed"},
		{name: "altered source", path: "parser.go", content: "arbitrary change", message: "original attribution/source changed"},
		{name: "altered manifest", path: "UPSTREAM-SHA256SUMS", content: "", message: "original manifest changed"},
		{name: "unaccounted owned source", path: "future.go", content: "package future", message: "unaccounted file"},
		{name: "original claimed as owned", path: "OWNED-FILES", content: "LICENSE\n", message: "overlapping owned"},
		{name: "missing owned file", path: "OWNED-FILES", content: "OWNED-FILES\nUPSTREAM-SHA256SUMS\nmissing.go\n", message: "missing inventory"},
		{name: "license maintenance prohibited", licenseException: true, message: "maintenance exception is not Go source/metadata"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			originals := map[string]string{"LICENSE": "original license", "COPYRIGHT": "original copyright", "parser.go": "original parser"}
			var manifest strings.Builder
			for _, name := range []string{"COPYRIGHT", "LICENSE", "parser.go"} {
				if err := os.WriteFile(filepath.Join(root, name), []byte(originals[name]), 0o644); err != nil {
					t.Fatal(err)
				}
				fmt.Fprintf(&manifest, "%x  %s\n", sha256.Sum256([]byte(originals[name])), name)
			}
			data := []byte(manifest.String())
			if err := os.WriteFile(filepath.Join(root, "UPSTREAM-SHA256SUMS"), data, 0o644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(root, "OWNED-FILES"), []byte("OWNED-FILES\nUPSTREAM-SHA256SUMS\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			p := policy{Manifests: map[string]string{"UPSTREAM-SHA256SUMS": digest(data)}}
			if tc.licenseException {
				p.Maintained = map[string]string{"LICENSE": digest([]byte("original license"))}
			}
			if tc.path != "" {
				path := filepath.Join(root, tc.path)
				var err error
				if tc.remove {
					err = os.Remove(path)
				} else {
					err = os.WriteFile(path, []byte(tc.content), 0o644)
				}
				if err != nil {
					t.Fatal(err)
				}
			}
			_, err := verify(root, p)
			if tc.message == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("expected %q, got %v", tc.message, err)
			}
		})
	}
}

func TestMissingOriginalPatterns(t *testing.T) {
	for _, name := range []string{
		"tools/benchmark/benchmark.go",
		".gitpod.yml",
		"syncers/auth_history_syncer/src/test/resources/impl/access.2022_06_09.log",
		"ui/keys/dev_x509_cert.cnf",
		"kubernetes/charts/athenz-zts/files/conf/athenz_conf.json",
		"original assets/file with spaces.txt",
	} {
		t.Run(name, func(t *testing.T) {
			root := t.TempDir()
			path := filepath.Join(root, name)
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				t.Fatal(err)
			}
			content := []byte("original public asset")
			if err := os.WriteFile(path, content, 0o644); err != nil {
				t.Fatal(err)
			}
			manifest := []byte(fmt.Sprintf("%x  %s\n", sha256.Sum256(content), name))
			if err := os.WriteFile(filepath.Join(root, "UPSTREAM-SHA256SUMS"), manifest, 0o644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(root, "OWNED-FILES"), []byte("OWNED-FILES\nUPSTREAM-SHA256SUMS\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			p := policy{Manifests: map[string]string{"UPSTREAM-SHA256SUMS": digest(manifest)}}
			if _, err := verify(root, p); err != nil {
				t.Fatalf("complete original inventory rejected: %v", err)
			}
			if err := os.Remove(path); err != nil {
				t.Fatal(err)
			}
			if _, err := verify(root, p); err == nil || !strings.Contains(err.Error(), "missing inventory files: "+name) {
				t.Fatalf("deleted ignored-pattern asset accepted: %v", err)
			}
		})
	}
}
