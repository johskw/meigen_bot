package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/router"
)

func init() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	router.SetRoutes(r)
	http.Handle("/", r)
}
