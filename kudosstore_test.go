package main

import (
	"testing"
)

func indexOf(list []string, search string) int {
	for index, item := range list {
		if item == search {
			return index
		}
	}
	return -1
}

const TestUser = "test-user"

func TestKudos(t *testing.T) {
	kudos := NewKudosStore()

	if kudos.Score(TestUser) != 0 {
		t.Fail()
	}

	if kudos.IncrBy(TestUser, 5) != 5 {
		t.Fail()
	}

	if kudos.Score(TestUser) != 5 {
		t.Fail()
	}

	list := kudos.Rankings()

	if indexOf(list, TestUser) < 0 {
		t.Fail()
	}

	kudos.Remove(TestUser)
}
