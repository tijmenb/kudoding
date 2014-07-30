package main

import (
	"encoding/json"
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
	Period string
}

func NewKudosStore() KudosStore {
	redisgo := os.Getenv("REDISTOGO_URL")
	client, err := redis.DialTimeout("tcp", redisgo, time.Duration(10)*time.Second)
	exitOnError(err)
	return KudosStore{Client: client}
}

func (store *KudosStore) SetPeriod(time time.Time) {
	store.Period = time.Format("2006/01")
}

func (store *KudosStore) kudosSet() string {
	return "kudos/" + store.Period
}

func (store *KudosStore) Score(name string) int {
	reply := store.Client.Cmd("zscore", store.kudosSet(), name)
	if reply.Type == redis.NilReply {
		return 0
	}
	kudos, err := store.Client.Cmd("zscore", store.kudosSet(), name).Int()
	exitOnError(err)
	return kudos
}

func (store *KudosStore) IncrBy(name string, kudos int) int {
	kudos, err := store.Client.Cmd("zincrby", store.kudosSet(), kudos, name).Int()
	exitOnError(err)
	return kudos
}

func (store *KudosStore) Rankings() []string {
	list, err := store.Client.Cmd("zrevrange", store.kudosSet(), 0, -1, "withscores").List()
	exitOnError(err)
	return list
}

func (store *KudosStore) Remove(name string) {
	reply := store.Client.Cmd("zrem", store.kudosSet(), name)
	exitOnError(reply.Err)
}

func (store *KudosStore) Del() {
	reply := store.Client.Cmd("del", store.kudosSet())
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

func (store *KudosStore) kudosLog(id string) string {
	return "kudoslog/" + store.Period + "/" + id
}

func (store *KudosStore) StoreKudos(ids []string, text string) {
	for _, id := range ids {
		reply := store.Client.Cmd("rpush", store.kudosLog(id), text)
		exitOnError(reply.Err)
	}
}

func (store *KudosStore) FetchKudos(id string) []string {
	list, err := store.Client.Cmd("lrange", store.kudosLog(id), 0, -1).List()
	exitOnError(err)
	return list
}
