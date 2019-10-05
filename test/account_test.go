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
  kind1 := model.NewKindOfItem(5, "オーブ")
  kind2 := model.NewKindOfItem(4, "剣")
  kind3 := model.NewKindOfItem(3, "服")
  alice.Items = []model.Item {
    model.NewItem (
      kind1,
    ),
    model.NewItem (
      kind2,
    ),
    model.NewItem (
      kind3,
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
  kind1 := model.NewKindOfItem(5, "オーブ")
  kind2 := model.NewKindOfItem(4, "剣")
  alice.Items = []model.Item {
    model.NewItem(kind1),
    model.NewItem(kind2),
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

func TestLogin(t *testing.T) {
  account := model.Login("player1", "password1")
  if account == nil {
    t.Error("ログインができなかった")
  }
  t.Log("Login終了")
}

func TestGetAccount(t *testing.T) {
  account := model.GetAccount("player1")
  if account == nil {
    t.Error("アカウントが取得できなかった")
  }
  t.Log("GetAccount終了")
}
func TestAccountsGiveItem(t *testing.T) {
  sender := model.GetAccount("player1")
  receiver := model.GetAccount("player2")
  item := sender.Items[0]
  senderItemLen := len(sender.Items)
  receiverItemLen := len(receiver.Items)
  success := model.GiveItem("player1", "player2", item.Id)
  if !success {
    t.Error("アイテム送信に失敗しました")
  }
  senderAfter := model.GetAccount("player1")
  receiverAfter := model.GetAccount("player2")
  if senderItemLen - 1 != len(senderAfter.Items) {
    t.Error("アイテムが適切に減っていない")
  }
  if receiverItemLen + 1 != len(receiverAfter.Items) {
    t.Error("アイテムが適切に増えていない")
  }
  if receiverAfter.Items[len(receiverAfter.Items) - 1].Id != item.Id {
    t.Error("適切なアイテムが増えていない")
  }
  t.Log("GiveItem終了")
}
