package model

import (
  "math"
)

var NewestItemId = 0
var NewestItemKindId = 0

type Item struct {
  Id int
  Rarity int
  Name string
  KindId int
}

type KindOfItem struct {
  Id int
  Rarity int
  Name string
  Price int
}

func NewKindOfItem(rarity int, name string) *KindOfItem {
  NewestItemKindId += 1
  return &KindOfItem {
    Id: NewestItemKindId,
    Rarity: rarity,
    Name: name,
    Price: int(math.Pow(10, float64(rarity))),
  }
}
func NewItem(kind *KindOfItem) Item {
  NewestItemId = NewestItemId + 1
  return Item {
    Id: NewestItemId,
    Rarity: kind.Rarity,
    Name: kind.Name,
    KindId: kind.Id,
  }
}

var KindsOfItem []*KindOfItem = []*KindOfItem {}

func SetKindOfItems (kindsOfItem []*KindOfItem) {
  KindsOfItem = kindsOfItem
}

func GetKindOfItem(id int) *KindOfItem {
  for _, kind := range KindsOfItem {
    if kind.Id == id {
      return kind
    }
  }
  return nil
}
