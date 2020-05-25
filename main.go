package main

import (
	"github.com/heroku/go-getting-started/src/apps/paperWiki"
	"github.com/heroku/go-getting-started/src/packages/core/router"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	router, port := router.GetRouter()

	router = paperWiki.GetRoutes(router)

	router.Run(":" + port)
}