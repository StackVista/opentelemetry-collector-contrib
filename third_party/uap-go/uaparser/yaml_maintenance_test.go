package uaparser

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDatabaseLoadingContracts(t *testing.T) {
	data := []byte(`user_agent_parsers:
  - regex: '(Fixture)/(\d+)'
    family_replacement: yes
    v1_replacement: 12
os_parsers:
  - regex: '(Fixture)/(\d+)'
    os_replacement: 'on'
    os_v1_replacement: '01'
device_parsers:
  - regex: '(Fixture)/(\d+)'
    device_replacement: 'Fixture Device'
`)
	p, err := NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	c := p.Parse("Fixture/7")
	if c.UserAgent.Family != "yes" || c.UserAgent.Major != "12" || c.Os.Family != "on" || c.Os.Major != "01" || c.Device.Family != "Fixture Device" {
		t.Fatalf("scalar/regex contract: %#v %#v %#v", c.UserAgent, c.Os, c.Device)
	}
	path := filepath.Join(t.TempDir(), "regexes.yaml")
	if err := ioutil.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	fileParser, err := New(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c, fileParser.Parse("Fixture/7")) {
		t.Fatal("file and byte loading differ")
	}
	optionsParser, err := NewWithOptions(path, EOsLookUpMode|EUserAgentLookUpMode|EDeviceLookUpMode, 100, 20, false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c, optionsParser.Parse("Fixture/7")) {
		t.Fatal("options loading differs")
	}
	if _, err := New(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("missing file must fail")
	}
	if _, err := NewFromBytes([]byte("user_agent_parsers: [")); err == nil {
		t.Fatal("malformed YAML must fail")
	}
	if _, err := NewFromBytes([]byte("user_agent_parsers: unexpected")); err == nil {
		t.Fatal("invalid database shape must fail")
	}
	if NewFromSaved().Parse("Mozilla/5.0").UserAgent == nil {
		t.Fatal("embedded database did not load")
	}
}
