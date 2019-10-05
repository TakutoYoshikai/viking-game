package model

var NewestItemId = 0

type Item struct {
  Id int
  Rarity int
  Name string
}

func NewItem(rarity int, name string) Item {
  NewestItemId = NewestItemId + 1
  return Item {
    Id: NewestItemId,
    Rarity: rarity,
    Name: name,
  }
}
func CopyItem(item *Item) Item{
  return NewItem(item.Rarity, item.Name)
}

var item1 Item = NewItem(
  1,
  "薬草",
)

var item2 Item = NewItem(
  2,
  "銅の鎧",
)

var item3 Item = NewItem(
  3,
  "金の盾",
)

var item4 Item = NewItem(
  4,
  "ダイヤの靴",
)

var item5 Item = NewItem(
  5,
  "伝説の剣",
)
var KindOfItems []*Item = []*Item {
  &item1,
  &item2,
  &item3,
  &item4,
  &item5,
}
