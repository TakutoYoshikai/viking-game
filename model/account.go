package model

import (
	"math/rand"
	"strconv"
	"time"
)

type Account struct {
	Username string
	Password string
	Items    []Item
}

type Accounts map[string]*Account

var accounts Accounts = Seeds()

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

func GetAccount(username string) *Account {
	return accounts[username]
}

func Login(username string, password string) *Account {
	account := GetAccount(username)
	if account.Password == password {
		return account
	}
	return nil
}

func GiveItem(from string, to string, itemId int) bool {
	fromAccount := GetAccount(from)
	toAccount := GetAccount(to)
	if fromAccount == nil || toAccount == nil {
		return false
	}
	return fromAccount.GiveItem(itemId, toAccount)
}

func chooseRandomItems() []Item {

	rand.Seed(time.Now().UnixNano())
	result := []Item{}
	if rand.Int63()%10 == 0 {
		result = append(result, NewItem(KindsOfItem[4]))
	}
	if rand.Int63()%5 == 0 {
		result = append(result, NewItem(KindsOfItem[3]))
	}
	if rand.Int63()%3 == 0 {
		result = append(result, NewItem(KindsOfItem[2]))
	}
	if rand.Int63()%2 == 0 {
		result = append(result, NewItem(KindsOfItem[1]))
	}
	result = append(result, NewItem(KindsOfItem[0]))
	return result
}
func Seeds() Accounts {
	var accounts Accounts = Accounts{}
	kindsOfItem := []*KindOfItem{
		NewKindOfItem(
			1,
			"木",
		),
		NewKindOfItem(
			2,
			"銅",
		),
		NewKindOfItem(
			3,
			"金",
		),
		NewKindOfItem(
			4,
			"ダイヤ",
		),
		NewKindOfItem(
			5,
			"伝説の剣",
		),
	}
	SetKindOfItems(kindsOfItem)
	for i := 0; i < 10000; i++ {
		accounts["player"+strconv.Itoa(i)] = &Account{
			Username: "player" + strconv.Itoa(i),
			Password: "password" + strconv.Itoa(i),
			Items:    chooseRandomItems(),
		}
	}
	accounts["rmt"] = &Account{
		Username: "rmt",
		Password: "rmt",
		Items:    chooseRandomItems(),
	}
	for i := 0; i < 10000; i++ {
		accounts["val"+strconv.Itoa(i)] = &Account{
			Username: "val" + strconv.Itoa(i),
			Password: "password" + strconv.Itoa(i),
			Items:    []Item{},
		}
	}
	accounts["val-clear"] = &Account{
		Username: "val-clear",
		Password: "clear",
		Items:    []Item{},
	}
	return accounts
}
