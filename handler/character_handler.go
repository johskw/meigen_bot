package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/model"
)

func PostCharacter(c *gin.Context) {
	var character model.Character
	err := c.Bind(&character)
	if err != nil {
		log.Print(err)
	}
	character, err = character.Create()
	if err != nil {
		log.Print(err)
	}
	nickname := model.Nickname{
		Nickname:    character.Name,
		CharacterID: character.ID,
	}
	_, err = nickname.Create()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func DeteleCharacter(c *gin.Context) {
	characterID, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteCharacter(characterID)
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
