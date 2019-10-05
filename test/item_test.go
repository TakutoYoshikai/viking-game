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

func TestCopyItem(t *testing.T) {
  item := model.NewItem(5, "sword")
  cpItem := model.CopyItem(&item)
  if item.Id >= cpItem.Id {
    t.Error("itemのコピーでidが不正")
  }
  if item.Name != cpItem.Name {
    t.Error("itemのNameがコピーできていない")
  }
  if item.Rarity != cpItem.Rarity {
    t.Error("itemのRarityがコピーできていない")
  }
}
