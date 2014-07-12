package web

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"net/url"
	"os"
	"testing"
	"time"
)

func ParseRedistogoUrl() (string, string) {
	redisUrl := os.Getenv("REDISTOGO_URL")
	redisInfo, _ := url.Parse(redisUrl)
	server := redisInfo.Host
	password := ""
	if redisInfo.User != nil {
		password, _ = redisInfo.User.Password()
	}
	return server, password
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err)
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

	kudos.Remove(TestUser)

	list := kudos.Rankings()
	fmt.Printf("%s\n", list)
}
