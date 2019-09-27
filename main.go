package main

import (
  "viking-game/model"
  "fmt"
  "github.com/gin-gonic/gin"
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
  router := gin.Default()
  router.GET("/:username/:password", func(ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := FindAccount(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    ctx.JSON(200, account.Items)
  })
  router.Run(":8080")
}

func main() {
  ConfigureServer()
}
