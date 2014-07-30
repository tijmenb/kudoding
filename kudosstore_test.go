package main

import (
	"testing"
	"time"
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

	time := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	kudos.SetPeriod(time)

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
