package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/handler"
)

func main() {
	r := gin.New()
	r.POST("/callback", handler.CallbackHandler)
	r.Run(":8080")
}
