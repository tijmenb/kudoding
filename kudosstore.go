package main

import (
	"encoding/json"
	"fmt"
	"github.com/fzzy/radix/redis"
	"log"
	"os"
	"time"
)

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

type KudosStore struct {
	Client *redis.Client
}

func NewKudosStore() KudosStore {
	client, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	exitOnError(err)
	return KudosStore{Client: client}
}

const KudosSet = "kudos"

func (store *KudosStore) Score(name string) int {
	reply := store.Client.Cmd("zscore", KudosSet, name)
	if reply.Type == redis.NilReply {
		return 0
	}
	kudos, err := store.Client.Cmd("zscore", KudosSet, name).Int()
	exitOnError(err)
	return kudos
}

func (store *KudosStore) IncrBy(name string, kudos int) int {
	kudos, err := store.Client.Cmd("zincrby", KudosSet, kudos, name).Int()
	exitOnError(err)
	return kudos
}

func (store *KudosStore) Rankings() []string {
	list, err := store.Client.Cmd("zrevrange", KudosSet, 0, -1, "withscores").List()
	exitOnError(err)
	return list
}

func (store *KudosStore) Remove(name string) {
	reply := store.Client.Cmd("zrem", KudosSet, name)
	exitOnError(reply.Err)
}

func (store *KudosStore) Del() {
	reply := store.Client.Cmd("del", KudosSet)
	exitOnError(reply.Err)
}

const UsersHash = "users"

func (store *KudosStore) StoreUsers(users []User) {
	reply := store.Client.Cmd("del", UsersHash)
	exitOnError(reply.Err)
	var hmset []interface{}
	for _, user := range users {
		jsonBlob, err := json.Marshal(user)
		exitOnError(err)
		hmset = append(hmset, "@"+user.Id)
		hmset = append(hmset, jsonBlob)
	}
	reply = store.Client.Cmd("hmset", UsersHash, hmset)
	exitOnError(reply.Err)
}

func (store *KudosStore) GetUsers(ids []string) []*User {
	if len(ids) == 0 {
		empty := []*User{}
		return empty
	}
	list, err := store.Client.Cmd("hmget", UsersHash, ids).List()
	exitOnError(err)
	users := []*User{}
	for _, jsonBlob := range list {
		user := User{}
		err := json.Unmarshal([]byte(jsonBlob), &user)
		if err == nil {
			users = append(users, &user)
		} else {
			users = append(users, nil)
		}
	}
	return users
}

const KudosLog = "kudoslog/%s"

func (store *KudosStore) StoreKudos(ids []string, text string) {
	for _, id := range ids {
		reply := store.Client.Cmd("rpush", fmt.Sprintf(KudosLog, id), text)
		exitOnError(reply.Err)
	}
}

func (store *KudosStore) FetchKudos(id string) []string {
	list, err := store.Client.Cmd("lrange", fmt.Sprintf(KudosLog, id), 0, -1).List()
	exitOnError(err)
	return list
}
