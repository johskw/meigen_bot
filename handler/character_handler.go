package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/model"
)

func CreateCharacter(c *gin.Context) {
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

func UpdateCharacter(c *gin.Context) {
	characterID, _ := strconv.Atoi(c.Param("id"))
	character, err := model.GetCharacter(characterID)
	if err != nil {
		log.Print(err)
	}
	var nickname model.Nickname
	for _, n := range character.Nicknames {
		if n.Nickname == character.Name {
			nickname = n
		}
	}
	err = c.Bind(&character)
	if err != nil {
		log.Print(err)
	}
	nickname.Nickname = character.Name
	_, err = character.Update()
	if err != nil {
		log.Print(err)
	}
	_, err = nickname.Update()
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
