package main

import (
  "github.com/gin-gonic/gin"
)
func CreateServer() *gin.Engine {
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
  return router
}
