package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/model"
)

func PostMeigen(c *gin.Context) {
	var meigen model.Meigen
	err := c.Bind(&meigen)
	if err != nil {
		log.Print(err)
	}
	meigen.CharacterID, _ = strconv.Atoi(c.Param("character_id"))
	_, err = meigen.Create()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/characters/"+c.Param("character_id"))
}
