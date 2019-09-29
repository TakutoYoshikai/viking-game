package seeds
import (
  "viking-game/model"
  "strconv"
  "math/rand"
  "time"
)
var item1 model.Item = model.NewItem(1, "薬草")
var item2 model.Item = model.NewItem(2, "銅の鎧")
var item3 model.Item = model.NewItem(3, "金の盾")
var item4 model.Item = model.NewItem(4, "ダイヤの靴")
var item5 model.Item = model.NewItem(5, "伝説の剣")

func chooseRandomItems() []model.Item {
  rand.Seed(time.Now().UnixNano())
  result := []model.Item {}
  if rand.Int63() % 10 == 0 {
    result = append(result, model.CopyItem(item5))
  }
  if rand.Int63() % 5 == 0 {
    result = append(result, model.CopyItem(item4))
  }
  if rand.Int63() % 3 == 0 {
    result = append(result, model.CopyItem(item3))
  }
  if rand.Int63() % 2 == 0 {
    result = append(result, model.CopyItem(item2))
  }
  result = append(result, model.CopyItem(item1))
  return result
}

func CreateAccounts() []model.Account {

  var accounts []model.Account = []model.Account {}
  for i := 0; i < 100; i++ {
    accounts = append(accounts, model.Account {
      Username: "player" + strconv.Itoa(i),
      Password: "password" + strconv.Itoa(i),
      Items: chooseRandomItems(),
    })
  }
  accounts = append(accounts, model.Account {
    Username: "rmt",
    Password: "rmt",
    Items: chooseRandomItems(),
  })
  return accounts
}
