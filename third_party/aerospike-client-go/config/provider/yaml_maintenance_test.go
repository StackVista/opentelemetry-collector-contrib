// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package provider_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aerospike/aerospike-client-go/v8/config/provider"
)

func TestYAMLLoadingContracts(t *testing.T) {
	for _, tc := range []struct {
		name, data string
		valid      bool
	}{
		{"valid", "version: '1'\ndynamic:\n  client:\n    rack_aware: true\n    app_id: '01'\n    timeout: 12\n", true},
		{"missing version", "dynamic: {}\n", false},
		{"malformed", "version: '1'\ndynamic: [\n", false},
		{"wrong shape", "version: '1'\ndynamic: scalar\n", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(p, []byte(tc.data), 0o600); err != nil {
				t.Fatal(err)
			}
			loader := provider.NewYamlConfigProviderWithPath(p)
			c := loader.LoadConfig(p)
			if (c != nil) != tc.valid {
				t.Fatalf("expected valid=%v, got %#v", tc.valid, c)
			}
			if c != nil {
				if *c.Version != "1" || !*c.Dynamic.Client.RackAware || *c.Dynamic.Client.ApplicationId != "01" || *c.Dynamic.Client.Timeout != 12 {
					t.Fatalf("scalar contract: %#v", c)
				}
				if loader.LoadConfig(p) != nil {
					t.Fatal("unchanged file must not reload")
				}
			}
		})
	}
	if provider.NewYamlConfigProvider().LoadConfig(filepath.Join(t.TempDir(), "missing.yaml")) != nil {
		t.Fatal("missing file must return nil")
	}
}
