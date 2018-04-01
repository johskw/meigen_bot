package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/handler"
)

func init() {
	r := gin.New()
	r.POST("/callback", handler.CallbackHandler)
	http.Handle("/", r)
}
