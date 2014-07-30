package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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

func giveKudos(users []*User, amount int) string {
	names := userNames(users)
	responses := []string{fmt.Sprintf("K U D O S given to %s", fancyJoin(names))}
	for _, name := range names {
		kudos.IncrBy(name, amount)
		responses = append(responses, fmt.Sprintf("%s has ranking %d", name, kudos.Score(name)))
	}
	return strings.Join(responses, "\n")
}

func retrieveKudos(users []*User) string {
	responses := []string{fmt.Sprintf("K U D O S given to\n")}
	for _, user := range users {
		list := kudos.FetchKudos("@" + user.Id)
		report := "\n  " + strings.Join(list, "\n  ")
		if len(list) == 0 {
			report = "none yet..."
		}
		responses = append(responses, fmt.Sprintf("@%s %s", user.Name, report))
	}
	return strings.Join(responses, "\n")
}

func ranking() string {
	ranks := kudos.Rankings()
	if len(ranks) == 0 {
		return "No K U D O S given yet...\nStart collecting and giving :-)"
	}
	responses := []string{"K U D O S ranking:"}
	for i := 0; i < len(ranks); i += 2 {
		responses = append(responses, fmt.Sprintf("%s has score %s", ranks[i], ranks[i+1]))
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
		return kudos.GetUsers(ids)
	}
	return users
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

func replaceUserNames(text string, users []*User) string {
	for _, user := range users {
		text = strings.Replace(text, "@"+user.Id, "@"+user.Name, 1)
	}
	return text
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
	kudos.SetPeriod(time.Now())

	w.Header().Set("Content-Type", "application/json")

	text := r.FormValue("text")
	user_id := r.FormValue("user_id")

	// first id and user is the from user
	ids := append([]string{"@" + user_id}, parseNames(text)...)
	users := handleUsers(ids)

	var answer string
	if hasNils(users) {
		answer = reportUnkownUsers(ids, users)
	} else if len(users) > 1 {
		amount := parseAmount(text)
		if amount != 0 {
			answer = giveKudos(users[1:], amount)
			now := time.Now().Format("Jan 2 at 3:04pm")
			text = text + " (from @" + user_id + " on " + now + ")"
			text = replaceUserNames(text, users)
			kudos.StoreKudos(ids[1:], text)
		} else {
			answer = retrieveKudos(users[1:])
		}
	} else {
		answer = ranking()
	}

	enc := json.NewEncoder(w)
	slackResponse := map[string]string{"text": answer}
	enc.Encode(slackResponse)
}
