
package test

import (
  "testing"
  "viking-game/accounts"
)

func TestLogin(t *testing.T) {
  account := accounts.Login("player1", "password1")
  if account == nil {
    t.Error("ログインができなかった")
  }
  t.Log("Login終了")
}

func TestGetAccount(t *testing.T) {
  account := accounts.GetAccount("player1")
  if account == nil {
    t.Error("アカウントが取得できなかった")
  }
  t.Log("GetAccount終了")
}

func TestAccountsGiveItem(t *testing.T) {
  sender := accounts.GetAccount("player1")
  receiver := accounts.GetAccount("player2")
  item := sender.Items[0]
  success := accounts.GiveItem("player1", "player2", item.Id)
  if !success {
    t.Error("アイテム送信に失敗しました")
  }
  senderAfter := accounts.GetAccount("player1")
  receiverAfter := accounts.GetAccount("player2")
  if len(sender.Items) - 1 != len(senderAfter.Items) {
    t.Error("アイテムが適切に減っていない")
  }
  if len(receiver.Items) + 1 != len(receiverAfter.Items) {
    t.Error("アイテムが適切に増えていない")
  }
  if receiverAfter.Items[len(receiverAfter.Items) - 1].Id != item.Id {
    t.Error("適切なアイテムが増えていない")
  }
  t.Log("GiveItem終了")
}
