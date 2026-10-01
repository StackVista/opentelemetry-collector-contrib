// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package yamlmatcher

import (
	"strings"
	"testing"

	"github.com/onsi/gomega/matchers"
)

func TestMatcherContract(t *testing.T) {
	m := &matchers.MatchYAMLMatcher{YAMLToMatch: "a: 1\nb: 'yes'"}
	ok, e := m.Match([]byte("b: yes\na: 1"))
	if e != nil || !ok {
		t.Fatalf("%v %v", ok, e)
	}
	if !strings.Contains(m.NegatedFailureMessage("a: 1\nb: yes"), "not to match YAML") {
		t.Fatal("missing negated diagnostic")
	}
	ok, e = m.Match("a: 2\nb: yes")
	if e != nil || ok || !strings.Contains(m.FailureMessage("a: 2\nb: yes"), "to match YAML") {
		t.Fatal("missing mismatch diagnostic")
	}
	for _, v := range []any{123, "a: ["} {
		if _, e = m.Match(v); e == nil {
			t.Fatalf("accepted %v", v)
		}
	}
}
