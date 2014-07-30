package main

import (
	"testing"
)

func TestFancyJoin(t *testing.T) {
	names := []string{"@me", "@myself", "@i"}
	fancy := fancyJoin(names)
	if fancy != "@me, @myself and @i" {
		t.Fail()
	}
}
