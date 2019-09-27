package main

import (
  "github.com/gin-gonic/gin"
  "strconv"
)
func CreateServer() *gin.Engine {
  router := gin.Default()
  router.GET("/:username/:password", func(ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    ctx.JSON(200, account.Items)
  })

  router.GET("/:username/:password/:itemid/:to",  func(ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    itemIdStr := ctx.Param("itemid")
    itemId, err := strconv.Atoi(itemIdStr)
    toId := ctx.Param("to")
    to := GetAccount(toId)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    account := Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    if to == nil {
      ctx.JSON(404, nil)
      return
    }
    success := GiveItem(account.Username, to.Username, itemId)
    if success {
      ctx.JSON(200, nil)
      return
    }
    ctx.JSON(500, nil)
  })
  return router
}
