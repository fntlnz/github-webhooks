package configuration

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	c := new(Configuration)
	c.ParseFile("../resources/test-configuration.json")

	if c.Host != "0.0.0.0" {
		t.Errorf("Unexpected configuration host, expected: `0.0.0.0` found: `%s`", c.Host)
	}

	if c.Port != "3091" {
		t.Errorf("Unexpected configuration port, expected: `3091` found: `%s`", c.Port)
	}

	if c.Path != "/usr/bin:/bin" {
		t.Errorf("Unexpected configuration path, expected: `/usr/bin:/bin` found: `%s`", c.Path)
	}

	if len(c.Repositories) != 1 {
		t.Errorf("Expected 1 configured repository, found %v", len(c.Repositories))
	}

	if _, ok := c.Repositories["fntlnz/dockerfiles"]; !ok {
		t.Errorf("Expected repository `fntlnz/dockerfiles` not found")
	}

	repo := c.Repositories["fntlnz/dockerfiles"]

	if _, ok := repo.Events["ping"]; !ok {
		t.Errorf("Expected event `ping` not found")
	}

	ping := repo.Events["ping"]

	if ping[0] != "touch ping-on-any-branch.txt" {
		t.Errorf("Unexpected command on event `ping` expected: `touch ping-on-any-branch.txt` found: `%v`", ping[0])
	}
}

func TestGetAddress(t *testing.T) {
	c := new(Configuration)
	c.ParseFile("../resources/test-configuration.json")
	address := c.GetAddress()
	expected := "0.0.0.0:3091"
	if address != expected {
		t.Errorf("Unexpected address generated: `%s`, expected %s", c.Port, expected)
	}
}

func TestGetAddressWithoutHost(t *testing.T) {
	c := new(Configuration)
	confJson := "{\"port\": \"3091\"}"
	c.Parse([]byte(confJson))
	address := c.GetAddress()
	expected := ":3091"
	if address != expected {
		t.Errorf("Unexpected address generated: `%s`, expected %s", c.Port, expected)
	}
}
