package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var kudos = NewKudosStore()

func init() {
	http.HandleFunc("/api", answerSlack)
}

func fancyJoin(parts []string) string {
	last := len(parts) - 1
	if last == 0 {
		return parts[last]
	}
	return strings.Join(parts[0:last], ", ") + " and " + parts[last]
}

func giveKudos(users []*User) string {
	names := userNames(users)
	responses := []string{fmt.Sprintf("K U D O S given to %s", fancyJoin(names))}
	for _, name := range names {
		kudos.IncrBy(name, 1)
		responses = append(responses, fmt.Sprintf("%s has ranking %d", name, kudos.Score(name)))
	}
	return strings.Join(responses, "\n")
}

func ranking() string {
	ranks := kudos.Rankings()
	if len(ranks) == 0 {
		return "No K U D O S given yet... lets get busy ;)\nhttps://www.youtube.com/watch?v=_gp51lt9kdA"
	}
	responses := []string{"K U D O S ranking:"}
	for i := 0; i < len(ranks); i += 2 {
		responses = append(responses, fmt.Sprintf("%s has ranking %s", ranks[i], ranks[i+1]))
	}
	return strings.Join(responses, "\n")
}

func refreshUserList() {
	users := UsersList()
	kudos.StoreUsers(users)
}

func handleUsers(ids []string) []*User {
	users := kudos.GetUsers(ids)
	if hasNils(users) {
		refreshUserList()
	}
	return kudos.GetUsers(ids)
}

func userNames(users []*User) []string {
	names := []string{}
	for _, user := range users {
		names = append(names, "@"+user.Name)
	}
	return names
}

func hasNils(list []*User) bool {
	for _, item := range list {
		if item == nil {
			return true
		}
	}
	return false
}

func reportUnkownUsers(ids []string, users []*User) string {
	responses := []string{}
	for index, user := range users {
		if user == nil {
			responses = append(responses, fmt.Sprintf("User %s is not known...", ids[index]))
		}
	}
	return strings.Join(responses, "\n")
}

func answerSlack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	text := r.FormValue("text")
	ids := parseNames(text)
	users := handleUsers(ids)

	var answer string
	if hasNils(users) {
		answer = reportUnkownUsers(ids, users)
	} else if len(users) > 0 {
		answer = giveKudos(users)
	} else {
		answer = ranking()
	}

	enc := json.NewEncoder(w)
	slackResponse := map[string]string{"text": answer}
	enc.Encode(slackResponse)
}
