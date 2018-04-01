package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/handler"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/callback", handler.CallbackHandler)

	r.LoadHTMLGlob("templates/*.tmpl")
	r.GET("/", handler.ShowIndexHandler)
	r.GET("/characters/new", handler.ShowCharacterForm)

	r.Run(":8080")
}
