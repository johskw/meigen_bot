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
	r.GET("/", handler.ShowIndex)
	r.GET("/characters/:id", handler.ShowCharacter)
	r.GET("/newcharacter", handler.ShowCharacterForm)
	r.POST("/createcharacter", handler.PostCharacter)

	r.POST("/createnickname/:character_id", handler.PostNickname)
	r.POST("/deletenickname/:nickname_id/:character_id", handler.DeleteNickname)

	r.POST("/createmeigen/:character_id", handler.PostMeigen)

	r.Run(":8080")
}
