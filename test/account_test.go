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
    model.Item {
      Rarity: 5,
      Name: "オーブ",
    },
    model.Item {
      Rarity: 4,
      Name: "剣",
    },
  }

  alice.GiveItem(0, bob)
  if len(alice.Items) != 1 {
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
