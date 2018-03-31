package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/johskw/meigen_bot/handler"
)

func main() {
	err := godotenv.Load("bot.env")
	if err != nil {
		panic(err)
	}
	r := gin.New()
	r.POST("/callback", handler.CallbackHandler)
	r.Run(":8080")
}
