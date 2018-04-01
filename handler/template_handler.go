package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/johskw/meigen_bot/model"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	characters, err := model.GetAllCharacters()
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"characters": characters,
	})
}

func ShowCharacterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "character_form.tmpl", nil)
}

func ShowCharacter(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	character, err := model.GetCharacter(id)
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "character.tmpl", gin.H{
		"character": character,
	})
}
