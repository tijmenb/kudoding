package main

import (
	"fmt"
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

func TestParseSlackTimeStamp(t *testing.T) {
	timestamp := parseSlackTimestamp("1355517523.000005")
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Sprintln(timestamp.Format(layout))
}
