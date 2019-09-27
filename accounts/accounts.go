
package accounts

import (
  "viking-game/model"
  "viking-game/seeds"
  "fmt"
)

var Accounts []model.Account = seeds.CreateAccounts()


func Login(username string, password string) *model.Account {
  for _, account := range Accounts {
    if account.Username == username && account.Password == password {
      return &account
    }
  }
  return nil
}

func GetAccount(username string) *model.Account {
  for _, account := range Accounts {
    if account.Username == username {
      return &account
    }
  }
  return nil
}

func getIndexOfAccount(username string) int {
  for i, account := range Accounts {
    if account.Username == username {
      return i
    }
  }
  return -1
}


func GiveItem(from string, to string, itemId int) bool {
  fromIndex := getIndexOfAccount(from)
  toIndex := getIndexOfAccount(to)
  if fromIndex == -1 || toIndex == -1 {
    return false
  }
  fromAccount := &Accounts[fromIndex]
  toAccount := &Accounts[toIndex]
  return fromAccount.GiveItem(itemId, toAccount)
}

func ShowAccounts() {
  fmt.Println(Accounts)
}
