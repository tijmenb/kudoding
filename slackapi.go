package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type User struct {
	Id       string
	Name     string
	RealName string `json:"real_name"`
	Profile  Profile
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Image     string `json:"image_48"`
}

type UserListResponse struct {
	Ok      bool
	Members []User
}

func UsersList() []User {
	token := os.Getenv("KUDOS_TOKEN")
	api_url := "https://slack.com/api/users.list"
	v := url.Values{}
	v.Set("token", token)

	res, err := http.Get(api_url + "?" + v.Encode())
	if err != nil {
		log.Fatal(err)
	}
	jsonBlob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userListResponse UserListResponse
	err = json.Unmarshal(jsonBlob, &userListResponse)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", userListResponse.Members)
	return userListResponse.Members
}
