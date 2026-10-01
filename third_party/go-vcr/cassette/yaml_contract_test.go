// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cassette_test

import (
	"github.com/dnaeon/go-vcr/cassette"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestYAMLCassetteContract(t *testing.T) {
	name := filepath.Join(t.TempDir(), "cassette")
	c := cassette.New(name)
	c.AddInteraction(&cassette.Interaction{Request: cassette.Request{Method: "GET", URL: "http://synthetic.invalid/fixture", Headers: http.Header{"X-Scalar": []string{"yes", "0123"}}}, Response: cassette.Response{Code: 200, Status: "200 OK", Body: "fixture"}})
	if e := c.Save(); e != nil {
		t.Fatal(e)
	}
	r, e := cassette.Load(name)
	if e != nil {
		t.Fatal(e)
	}
	req, _ := http.NewRequest("GET", "http://synthetic.invalid/fixture", nil)
	i, e := r.GetInteraction(req)
	if e != nil || i.Response.Body != "fixture" || i.Request.Headers.Get("X-Scalar") != "yes" {
		t.Fatalf("%+v %v", i, e)
	}
	if _, e = r.GetInteraction(req); e == nil {
		t.Fatal("unexpected replay twice")
	}
	r.ReplayableInteractions = true
	if _, e = r.GetInteraction(req); e != nil {
		t.Fatal(e)
	}
	raw, _ := os.ReadFile(name + ".yaml")
	if !strings.Contains(string(raw), "version: 1") {
		t.Fatalf("%s", raw)
	}
	if e = os.WriteFile(name+".yaml", []byte("interactions: ["), 0o600); e != nil {
		t.Fatal(e)
	}
	if _, e = cassette.Load(name); e == nil {
		t.Fatal("accepted malformed YAML")
	}
}
