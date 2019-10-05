package test

import (
  "testing"
  "viking-game/model"
  "math"
)

func TestNewItem(t *testing.T) {
  firstItemId := model.NewestItemId
  kind1 := model.NewKindOfItem(5, "オーブ")
  kind2 := model.NewKindOfItem(4, "剣")
  if kind1.Price != int(math.Pow(10, float64(kind1.Rarity))) {
    t.Error("itemのpriceが適切でない")
  }
  if kind2.Price != int(math.Pow(10, float64(kind2.Rarity))) {
    t.Error("itemのpriceが適切でない")
  }
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

