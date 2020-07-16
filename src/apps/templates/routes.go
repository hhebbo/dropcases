package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoutes(router *gin.Engine) *gin.Engine {
	router.GET(MAIN_URL, func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"templates_home.tmpl.html",
			gin.H{},
		)
	})

	router.GET(TEMPLATES_URL, func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"templates_templates.tmpl.html",
			gin.H{},
		)
	})

	router.GET(INTERNAL_TEMPLATES_URL, func(c *gin.Context) {
		CreateFoldersFromTemplate()
		CreateFoldersFromTemplate()

		// c.HTML(
		// 	http.StatusOK,
		// 	"templates_internal_template.tmpl.html",
		// 	gin.H{},
		// )
	})

	return router
}
