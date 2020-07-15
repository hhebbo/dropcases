package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/hhebbo/dropcases/src/apps/home"
	"github.com/hhebbo/dropcases/src/apps/paperWiki"
	"github.com/hhebbo/dropcases/src/apps/templates"
	"github.com/hhebbo/dropcases/src/packages/core/router"
)

func main() {
	router, port := router.GetRouter()

	router = home.GetRoutes(router)
	router = templates.GetRoutes(router)
	router = paperWiki.GetRoutes(router)

	router.Run(":" + port)
}
