package main

import (
  "net/http"
  "viking-game/model"
  "encoding/json"
)

var Accounts []model.Account


func FindAccount(username string, password string) *model.Account {
  for _, account := range Accounts {
    if account.Username == username && account.Password == password {
      return &account
    }
  }
  return nil
}

func ConfigureServer() {
  http.HandleFunc("/user", func (w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    password := r.URL.Query().Get("password")
    account := FindAccount(username, password)
    res, err := json.Marshal(account.Items)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
  })
}

func main() {
  ConfigureServer()
  http.ListenAndServe(":8080", nil)
}
