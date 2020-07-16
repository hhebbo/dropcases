package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hhebbo/dropcases/src/apps/appConfig"
)

func GetRouter() (*gin.Engine, string) {
	// port := config.GetValue(appConfig.DROPCASES_PORT)
	port := appConfig.DROPCASES_PORT
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.StaticFile("/img.jpg", "./templates/img.jpg")
	router.Static("/static", "static")

	return router, port
}
