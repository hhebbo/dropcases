package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/src/apps/appConfig"
	"github.com/heroku/go-getting-started/src/packages/core/config"
)

func GetRouter() (*gin.Engine, string) {
	port := config.GetValue(appConfig.DROPCASES_PORT)

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	return router, port
}
