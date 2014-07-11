package robokudos

import (
  "os"
  "fmt"
  "github.com/fzzy/radix/redis"
  "time"
  "testing"
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

// func ParseRedistogoUrl() (string, string) {
//   redisUrl := os.Getenv("REDISTOGO_URL")
//   redisInfo, _ := url.Parse(redisUrl)
//   server := redisInfo.Host
//   password := ""
//   if redisInfo.User != nil {
//     password, _ = redisInfo.User.Password()
//   }
//   return server, password
// }

func errHndlr(err error) {
  if err != nil {
    fmt.Println("error!!", err)
    os.Exit(1)
  }
}


func TestLoad(t *testing.T) {
  client, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
  errHndlr(err)

  r = client.Cmd("set", "mykey0", "myval0")
  errHndlr(r.Err)

  savedKey, err = client.Cmd("get", "mykey0").Str()
  errHndlr(err)

  fmt.Println("Err: ", err)
  fmt.Println("mykey0:", savedKey)

  // Load()
  // k := KudosType {Name: "Sander", Kudos: 0}
  // db := KudosDatabase {}
  // db.ApplyKudos("Sander", 1)
  // fmt.Printf("%s\n", k)
  // fmt.Printf("%s\n", db)
}
