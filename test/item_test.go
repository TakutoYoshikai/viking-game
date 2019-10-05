package test

import (
  "testing"
  "viking-game/model"
)

func TestNewItem(t *testing.T) {
  firstItemId := model.NewestItemId
  kind1 := model.NewKindOfItem(5, "オーブ")
  kind2 := model.NewKindOfItem(4, "剣")
  item1 := model.NewItem(kind1)
  item2 := model.NewItem(kind2)
  if item1.Id != firstItemId + 1 {
    t.Error("itemのidが足されていない")
  }
  if item2.Id != item1.Id + 1 {
    t.Error("itemのidが足されていない")
  }
  if item1.Name != "オーブ" {
    t.Error("itemのnameが反映されていない")
  }
  if item1.Rarity != 5 {
    t.Error("itemのrarityが反映されていない")
  }
  t.Log("NewItem終了")
}

