package model
import (
  "fmt"
)


type Account struct {
  Username string
  Password string
  Items []Item
}

func (from *Account) GiveItem(index int, to *Account) bool {
  if index >= len(from.Items) {
    fmt.Println("couldn't give an item")
    return false
  }
  item := from.Items[index]
  from.Items = append(from.Items[:index], from.Items[index+1:]...)
  to.Items = append(to.Items, item)
  return true
}
