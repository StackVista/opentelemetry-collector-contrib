// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestForeignOwnerCoverageIDs(t *testing.T) {
	root := t.TempDir()
	modules := map[string]string{
		"receiver/examplereceiver":          "example.com/collector/receiver/examplereceiver",
		"third_party/datadog-logs-config":   "github.com/DataDog/datadog-agent/comp/logs/agent/config",
		"third_party/datadog-nodetreemodel": "github.com/DataDog/datadog-agent/pkg/config/nodetreemodel",
	}
	for path, name := range modules {
		dir := filepath.Join(root, filepath.FromSlash(path))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module "+name+"\n"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	config, err := walkTree(Args{BasePrefix: "example.com/collector", Dir: root})
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]string{
		"receiver_example":                  "receiver/examplereceiver/**",
		"third_party_datadog-logs-config":   "third_party/datadog-logs-config/**",
		"third_party_datadog-nodetreemodel": "third_party/datadog-nodetreemodel/**",
	}
	for _, component := range config.ComponentManagement.IndividualComponents {
		path, exists := expected[component.ComponentID]
		if !exists || len(component.Paths) != 1 || filepath.ToSlash(component.Paths[0]) != path {
			t.Fatalf("unexpected component: %+v", component)
		}
		delete(expected, component.ComponentID)
	}
	if len(expected) != 0 {
		t.Fatalf("missing coverage: %v", expected)
	}
}
