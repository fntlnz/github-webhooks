package main

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	c := new(Configuration)
	c.Parse("resources/test-configuration.json")

	if c.Port != "3091" {
		t.Errorf("Unexpected configuration port, expected: `3091` found: `%v`", c.Port)
	}

	if c.Path != "/usr/bin:/bin" {
		t.Errorf("Unexpected configuration path, expected: `/usr/bin:/bin` found: `%s`", c.Path)
	}

	if len(c.Repositories) != 1 {
		t.Errorf("Expected 1 configured repository, found %v", len(c.Repositories))
	}

	if _, ok := c.Repositories["fntlnz/statik"]; !ok {
		t.Errorf("Expected repository `fntlnz/statik` not found")
	}


	repo := c.Repositories["fntlnz/statik"]

	if _, ok := repo.Events["ping"]; !ok {
		t.Errorf("Expected event `ping` not found")
	}

	ping := repo.Events["ping"]

	if ping[0] != "touch ping-on-any-branch.txt" {
		t.Errorf("Unexpected command on event `ping` expected: `touch ping-on-any-branch.txt` found: `%v`", ping[0])
	}

}
