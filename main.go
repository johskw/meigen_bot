package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/johskw/meigen_bot/router"
)

func main() {
	err := godotenv.Load("bot.env")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("app/templates/*.tmpl")
	router.SetRoutes(r)
	r.Run(":8080")
}
