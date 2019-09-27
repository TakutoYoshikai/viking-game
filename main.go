package main

import (
  "net/http"
  "viking-game/model"
  "encoding/json"
  "fmt"
)

var Accounts []model.Account = MakeSeed()


func FindAccount(username string, password string) *model.Account {
  for _, account := range Accounts {
    fmt.Println(account)
    if account.Username == username && account.Password == password {
      return &account
    }
  }
  return nil
}

func ConfigureServer() {
  http.HandleFunc("/user", func (w http.ResponseWriter, r *http.Request) {
    password := r.URL.Query().Get("password")
    username := r.URL.Query().Get("username")
    fmt.Println(username)
    fmt.Println(password)
    account := FindAccount(username, password)
    w.Header().Set("Content-Type", "application/json")
    if account == nil {
      http.Error(w, "Unauthorized", http.StatusUnauthorized)
      return
    }
    res, err := json.Marshal(account.Items)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Write(res)
  })
}

func main() {
  ConfigureServer()
  http.ListenAndServe(":8080", nil)
}
