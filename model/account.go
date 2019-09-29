package model

import (
  "math/rand"
  "time"
  "strconv"
)

type Account struct {
  Username string
  Password string
  Items []Item
}

type Accounts []*Account
var accounts Accounts = Seeds()

func delItem(items []Item, i int) []Item {
  items = append(items[:i], items[i+1:]...)
  return items
}

func (account *Account) SearchItemAndIndex(id int) (int, *Item) {
  for i, item := range account.Items {
    if item.Id == id {
      return i, &item
    }
  }
  return -1, nil
}
func (from *Account) GiveItem(id int, to *Account) bool {
  index, item := from.SearchItemAndIndex(id)
  if index == -1 {
    return false
  }
  from.Items = delItem(from.Items, index)
  to.Items = append(to.Items, *item)
  return true
}


func GetAccount(username string) *Account {
  for _, account := range accounts {
    if account.Username == username {
      return account
    }
  }
  return nil
}

func Login(username string, password string) *Account {
  account := GetAccount(username)
  if account.Password == password {
    return account
  }
  return nil
}

func GiveItem(from string, to string, itemId int) bool {
  fromAccount := GetAccount(from)
  toAccount := GetAccount(to)
  if fromAccount == nil || toAccount == nil {
    return false
  }
  return fromAccount.GiveItem(itemId, toAccount)
}

func chooseRandomItems() []Item {
  var item1 Item = NewItem(1, "薬草")
  var item2 Item = NewItem(2, "銅の鎧")
  var item3 Item = NewItem(3, "金の盾")
  var item4 Item = NewItem(4, "ダイヤの靴")
  var item5 Item = NewItem(5, "伝説の剣")

  rand.Seed(time.Now().UnixNano())
  result := []Item {}
  if rand.Int63() % 10 == 0 {
    result = append(result, CopyItem(item5))
  }
  if rand.Int63() % 5 == 0 {
    result = append(result, CopyItem(item4))
  }
  if rand.Int63() % 3 == 0 {
    result = append(result, CopyItem(item3))
  }
  if rand.Int63() % 2 == 0 {
    result = append(result, CopyItem(item2))
  }
  result = append(result, CopyItem(item1))
  return result
}
func Seeds() Accounts {
  var accounts Accounts = Accounts{}
  for i := 0; i < 100; i++ {
    accounts = append(accounts, &Account {
      Username: "player" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Items: chooseRandomItems(),
    })
  }
  accounts = append(accounts, &Account {
    Username: "rmt",
    Password: "rmt",
    Items: chooseRandomItems(),
  })
  return accounts
}
