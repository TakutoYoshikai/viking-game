package test

import (
  "testing"
  "viking-game/model"
)

func TestGiveItem(t *testing.T) {
  alice := &model.Account {
    Username: "alice",
    Password: "alice01",
  }

  bob := &model.Account {
    Username: "bob",
    Password: "bob01",
  }
  alice.Items = []model.Item {
    model.NewItem (
      5,
      "オーブ",
    ),
    model.NewItem (
      4,
      "剣",
    ),
    model.NewItem (
      3,
      "服",
    ),
  }

  alice.GiveItem(alice.Items[0].Id, bob)
  if len(alice.Items) != 2 {
    t.Error("aliceのitemが減っていない: ", len(alice.Items))
  }
  if len(bob.Items) != 1 {
    t.Error("bobのitemが増えていない: ", len(bob.Items)) 
  }
  if alice.Items[0].Name != "剣" {
    t.Error("aliceの残っているitemが間違っている: ", alice.Items[0].Name)
  }
  if bob.Items[0].Name != "オーブ" {
    t.Error("bobが受け取っているアイテムが違う: ", bob.Items[0].Name)
  }
  t.Log("GiveItem終了")
}

func TestSearchItemAndIndex(t *testing.T) {
  alice := &model.Account {
    Username: "alice",
    Password: "alice01",
  }
  alice.Items = []model.Item {
    model.NewItem(5, "オーブ"),
    model.NewItem(4, "剣"),
  }
  index, _ := alice.SearchItemAndIndex(alice.Items[1].Id + 1)
  if index != -1 {
    t.Error("所有していないitemのidで検索して、indexが-1でない")
  }
  index, item := alice.SearchItemAndIndex(alice.Items[1].Id)
  if index != 1 {
    t.Error("間違ったindexを返している")
  }
  if item.Id != alice.Items[1].Id {
    t.Error("間違ったitemを返している")
  }
  t.Log("SearchItemAndIndex終了")
}
