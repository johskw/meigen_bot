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
	r.GET("/character/:id", handler.ShowCharacter)
	r.GET("/newcharacter", handler.ShowNewCharacterForm)
	r.POST("/createcharacter", handler.PostCharacter)
	r.GET("/editcharacter/:id", handler.ShowEditCharacterForm)
	r.POST("/deletecharacter/:id", handler.DeteleCharacter)

	r.POST("/createnickname/:character_id", handler.PostNickname)
	r.POST("/deletenickname/:nickname_id/:character_id", handler.DeleteNickname)

	r.POST("/createmeigen/:character_id", handler.PostMeigen)
	r.POST("/deletemeigen/:meigen_id/:character_id", handler.DeleteMeigen)

	r.Run(":8080")
}
