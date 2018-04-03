package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/handler"
)

func SetRoutes(r *gin.Engine) {
	r.POST("/callback", handler.Callback)

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("AUTH_USER"): os.Getenv("AUTH_PATH"),
	}))

	auth.GET("/", handler.ShowIndex)
	auth.GET("/character/:id", handler.ShowCharacter)
	auth.GET("/newcharacter", handler.ShowNewCharacterForm)
	auth.POST("/createcharacter", handler.CreateCharacter)
	auth.GET("/editcharacter/:id", handler.ShowEditCharacterForm)
	auth.POST("/updatecharacter/:id", handler.UpdateCharacter)
	auth.POST("/deletecharacter/:id", handler.DeteleCharacter)

	auth.POST("/createnickname/:character_id", handler.CreateNickname)
	auth.POST("/deletenickname/:nickname_id/:character_id", handler.DeleteNickname)

	auth.POST("/createmeigen/:character_id", handler.CreateMeigen)
	auth.POST("/deletemeigen/:meigen_id/:character_id", handler.DeleteMeigen)

	auth.GET("/checker", handler.ShowChecker)
	auth.POST("/checker", handler.ShowResult)
}
