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
func CopyItem(item Item) Item{
  return NewItem(item.Rarity, item.Name)
}
