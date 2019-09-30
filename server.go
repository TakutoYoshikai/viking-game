package main

import (
  "github.com/gin-gonic/gin"
  "viking-game/model"
  "strconv"
)
func CreateServer() *gin.Engine {
  router := gin.Default()
  router.GET("/item/:username/:item_id", func(ctx *gin.Context) {
    username := ctx.Param("username")
    itemIdStr := ctx.Param("item_id")
    itemId, err := strconv.Atoi(itemIdStr)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }

    account := model.GetAccount(username)
    _, item := account.SearchItemAndIndex(itemId)
    if item == nil {
      ctx.JSON(404, nil)
      return
    }
    ctx.JSON(200, item)
  })
  router.GET("/items/:username", func(ctx *gin.Context) {
    username := ctx.Param("username")
    account := model.GetAccount(username)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    ctx.JSON(200, account.Items)
  })
  router.GET("/users/:username/:password", func(ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(403, nil)
      return
    }
    ctx.JSON(200, account)
  })

  router.GET("/send/:username/:password/:item_id/:to",  func(ctx *gin.Context) {
    username := ctx.Param("username")
    password := ctx.Param("password")
    itemIdStr := ctx.Param("item_id")
    itemId, err := strconv.Atoi(itemIdStr)
    toId := ctx.Param("to")
    to := model.GetAccount(toId)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    account := model.Login(username, password)
    if account == nil {
      ctx.JSON(401, nil)
      return
    }
    if to == nil {
      ctx.JSON(404, nil)
      return
    }
    success := model.GiveItem(account.Username, to.Username, itemId)
    if success {
      ctx.JSON(200, nil)
      return
    }
    ctx.JSON(500, nil)
  })
  return router
}
