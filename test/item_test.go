package test

import (
  "testing"
  "viking-game/model"
)

func TestNewItem(t *testing.T) {
  item1 := model.NewItem(5, "sword")
  item2 := model.NewItem(4, "hat")
  if item1.Id != 1 {
    t.Error("最初のitemのidが1ではない: ", item1.Id)
  }
  if item2.Id != 2 {
    t.Error("2番目のitemのidが2ではない: ", item2.Id)
  }
  t.Log("NewItem終了")
}
