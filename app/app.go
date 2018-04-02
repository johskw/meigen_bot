package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/handler"
)

func init() {
	r := gin.New()

	r.POST("/callback", handler.CallbackHandler)

	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", handler.ShowIndex)
	r.GET("/character/:id", handler.ShowCharacter)
	r.GET("/newcharacter", handler.ShowNewCharacterForm)
	r.POST("/createcharacter", handler.CreateCharacter)
	r.GET("/editcharacter/:id", handler.ShowEditCharacterForm)
	r.POST("/updatecharacter/:id", handler.UpdateCharacter)
	r.POST("/deletecharacter/:id", handler.DeteleCharacter)

	r.POST("/createnickname/:character_id", handler.CreateNickname)
	r.POST("/deletenickname/:nickname_id/:character_id", handler.DeleteNickname)

	r.POST("/createmeigen/:character_id", handler.CreateMeigen)
	r.POST("/deletemeigen/:meigen_id/:character_id", handler.DeleteMeigen)

	http.Handle("/", r)
}
