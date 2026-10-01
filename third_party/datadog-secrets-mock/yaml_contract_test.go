// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"strings"
	"testing"
)

func TestYAMLSerializationContract(t *testing.T) {
	m := New(t)
	m.SetSecrets(map[string]string{"handle": "0123"})
	calls := 0
	m.SubscribeToChanges(func(_ string, _ string, _ []string, _ any, _ any) { calls++ })
	b, e := m.Resolve([]byte("password: ENC[handle]\nflag: yes\nquoted: '01'\n"), "fixture", "", "")
	if e != nil || !strings.Contains(string(b), `password: "0123"`) || !strings.Contains(string(b), "flag: true") || !strings.Contains(string(b), `quoted: "01"`) || calls != 1 {
		t.Fatalf("%s %v calls=%d", b, e, calls)
	}
	for _, s := range []string{"password: ENC[missing]", "bad: ["} {
		if _, e = m.Resolve([]byte(s), "fixture", "", ""); e == nil {
			t.Fatalf("accepted %q", s)
		}
	}
}
