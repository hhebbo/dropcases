package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"home.tmpl.html",
			gin.H{},
		)
	})

	return router
}
