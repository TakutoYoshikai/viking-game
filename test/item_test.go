package test

import (
  "testing"
  "viking-game/model"
)

func TestNewItem(t *testing.T) {
  firstItemId := model.NewestItemId
  item1 := model.NewItem(5, "sword")
  item2 := model.NewItem(4, "hat")
  if item1.Id != firstItemId + 1 {
    t.Error("itemのidが足されていない")
  }
  if item2.Id != item1.Id + 1 {
    t.Error("itemのidが足されていない")
  }
  t.Log("NewItem終了")
}
