package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("app/templates/*.tmpl")
	router.SetRoutes(r)
	r.Run(":8080")
}
