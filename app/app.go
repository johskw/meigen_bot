package app

import (
	"net/http"

	"github.com/johskw/meigen_bot/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("bot.env")
	if err != nil {
		panic(err)
	}
	r := gin.New()
	r.POST("/callback", handler.CallbackHandler)
	http.Handle("/", r)
}
