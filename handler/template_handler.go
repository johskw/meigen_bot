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

func ShowNewCharacterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new_character.tmpl", nil)
}

func ShowEditCharacterForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	character, err := model.GetCharacter(id)
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "edit_character.tmpl", gin.H{
		"character": character,
	})
}

func ShowChecker(c *gin.Context) {
	c.HTML(http.StatusOK, "checker.tmpl", nil)
}

func ShowResult(c *gin.Context) {
	meigenText, err := model.GetMeigenFromNickname(c.PostForm("nickname"))
	if err != nil {
		log.Print(err)
	}
	c.HTML(http.StatusOK, "checker.tmpl", gin.H{
		"nickname": c.PostForm("nickname"),
		"meigen":   meigenText,
	})
}
