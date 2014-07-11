package robokudos

import (
 "fmt"
 //  "os"
 //  "io/ioutil"
 //  "encoding/json"
	"testing"
  "github.com/xuyu/goredis"
)

// type KudosType struct {
//   Name string
//   Kudos int
// }

// type KudosDatabase struct {
//   Kudos []KudosType
// }

// func (db KudosDataBase) Find(name string) int {
//   for index, kudos := range db.Kudos {
//     if kudos.Name == name {
//       return index
//     }
//   }
//   return -1
// }

// func (db KudosDataBase) ApplyKudos(name string, kudos int) {
//   index := db.Find(name)
//   if index >= 0 {
//     db.Kudos[index].Kudos += kudos
//   } else {
//     kudos := KudosType{Name: name, Kudos: kudos}
//     append(db.Kudos, kudos)
//   }
// }

// const KudosFilename = "./kudos.json"

// func Load() KudosDatabase {
// 	file, e := ioutil.ReadFile(KudosFilename)
// 	if e != nil {
// 		fmt.Printf("File error: %v\n", e)
// 		os.Exit(1)
// 	}
// }

// func Save(KudosDatabase) {
//   file, e := ioutil.WriteFile(KudosFilename)
//   if e != nil {
//     fmt.Printf("File error: %v\n", e)
//     os.Exit(1)
//   }
// }

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

func RedisConnect() {
  client, err := Dial()
  client, err := Dial(&DialConfig{Address: "127.0.0.1:6379"})
  client, err := DialURL("tcp://127.0.0.1:6379/0?timeout=10s&maxidle=1")
  if e != nil {
    fmt.Printf("File error: %v\n", e)
    os.Exit(1)
  }
}

// func SetKudos() {
//   err := client.Set("key", "value", 0, 0, false, false)
// }

// func GetKudos() {
//   value, err := client.Get("key")    
// }

func TestLoad(t *testing.T) {
  RedisConnect()
	// Load()
  // k := KudosType {Name: "Sander", Kudos: 0}
  // db := KudosDatabase {}
  // db.ApplyKudos("Sander", 1)
  // fmt.Printf("%s\n", k)
  // fmt.Printf("%s\n", db)
}
