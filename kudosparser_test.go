package main

import (
	"strings"
	"testing"
)

func stringsEquals(a, b []string) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestKudosParser(t *testing.T) {
	testNames := []string{"@me", "@myself", "@i"}
	names := parseNames(strings.Join(testNames, " and "))
	if !stringsEquals(names, testNames) {
		t.Fail()
	}
}
