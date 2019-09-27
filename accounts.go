
package main

import (
  "viking-game/model"
  "fmt"
)

var Accounts []model.Account = MakeSeed()


func FindAccount(username string, password string) *model.Account {
  for _, account := range Accounts {
    fmt.Println(account)
    if account.Username == username && account.Password == password {
      return &account
    }
  }
  return nil
}
