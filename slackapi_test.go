package main

import (
	"testing"
)

func TestUserList(t *testing.T) {
	apiUsers := UsersList()
	if len(apiUsers) < 1 {
		t.Fail()
	}
	kudos := NewKudosStore()
	kudos.StoreUsers(apiUsers)
	ids := []string{}
	for _, user := range apiUsers {
		ids = append(ids, "@"+user.Id)
	}
	ids = append(ids, "@unknown")

	users := kudos.GetUsers(ids)
	for index, user := range users {
		if user == nil && ids[index] == "@unknown" {
			continue
		} else if user.Name != users[index].Name {
			t.Fail()
		}
	}
}
