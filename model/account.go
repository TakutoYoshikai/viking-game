package model


type Account struct {
  Username string
  Password string
  Items []Item
}

func delItem(items []Item, i int) []Item {
  items = append(items[:i], items[i+1:]...)
  return items
}

func (account *Account) SearchItemAndIndex(id int) (int, *Item) {
  for i, item := range account.Items {
    if item.Id == id {
      return i, &item
    }
  }
  return -1, nil
}
func (from *Account) GiveItem(id int, to *Account) bool {
  index, item := from.SearchItemAndIndex(id)
  if index == -1 {
    return false
  }
  from.Items = delItem(from.Items, index)
  to.Items = append(to.Items, *item)
  return true
}
