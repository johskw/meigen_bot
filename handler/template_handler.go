package handler

import (
	"log"
	"net/http"

	"github.com/johskw/meigen_bot/model"

	"github.com/gin-gonic/gin"
)

func ShowIndexHandler(c *gin.Context) {
	characters, err := model.GetAllCharacters()
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"characters": characters,
	})
}
