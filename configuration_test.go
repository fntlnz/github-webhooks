package main

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	c := new(Configuration)
	c.Parse("github-webhooks.json")

	if c.Port != "3091" {
		t.Errorf("Unexpected configuration port, expected: 3091 found: %v", c.Port)
	}

	if len(c.Repositories) != 1 {
		t.Errorf("Expected 1 configured repository, found %v", len(c.Repositories))
	}

	if _, ok := c.Repositories["majinbuu/statik"]; !ok {
		t.Errorf("Expected repository `majinbuu/statik` not found")
	}


	repo := c.Repositories["majinbuu/statik"]

	if _, ok := repo.Events["ping"]; !ok {
		t.Errorf("Expected event `ping` not found")
	}

	ping := repo.Events["ping"]

	if ping[0] != "touch ping-on-any-branch.txt" {
		t.Errorf("Unexpected command on event `ping` expected: `touch ping-on-any-branch.txt` found: %v", ping[0])
	}

}
