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
		t.Error()
	}
}

func TestKudosAmount(t *testing.T) {
	if parseAmount(" no amount ") != 0 {
		t.Error()
	}
	if parseAmount(" +2 ") != 2 {
		t.Error()
	}
	if parseAmount(" -3 ") != -3 {
		t.Error()
	}
}
