// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package viper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSelectedConfigurationContracts(t *testing.T) {
	v := New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(strings.NewReader("enabled: yes\ncode: '01'\ncount: 12\n")); err != nil {
		t.Fatal(err)
	}
	if v.Get("enabled") != true || v.GetString("code") != "01" || v.GetInt("count") != 12 {
		t.Fatalf("scalar contract: %#v", v.AllSettings())
	}
	if err := v.ReadConfig(strings.NewReader("value: [")); err == nil {
		t.Fatal("malformed YAML accepted")
	}
	// This exact selected DataDog revision already excludes HCL from file loading.
	p := filepath.Join(t.TempDir(), "config.hcl")
	if err := os.WriteFile(p, []byte("value = 12\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	h := New()
	h.SetConfigFile(p)
	if err := h.ReadInConfig(); err != UnsupportedConfigError("hcl") {
		t.Fatalf("selected HCL file contract changed: %v", err)
	}
}
