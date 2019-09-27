package model


type Account struct {
  Username string
  Password string
  Items []Item
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
  index, _ := from.SearchItemAndIndex(id)
  if index == -1 {
    return false
  }
  item := from.Items[index]
  from.Items = append(from.Items[:index], from.Items[index+1:]...)
  to.Items = append(to.Items, item)
  return true
}
