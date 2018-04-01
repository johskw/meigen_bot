package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/model"
)

func PostCharacter(c *gin.Context) {
	var character model.Character
	err := c.Bind(&character)
	if err != nil {
		log.Print(err)
	}
	_, err = character.Create()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
