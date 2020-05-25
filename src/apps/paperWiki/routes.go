package paperWiki

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MAIN_URL       = "/paper-wiki"
	PAGE_URL       = "/page"
	ROOT_WIKI_PATH = "/Paper Wiki"
)

func getRoutes(router *gin.Engine) *gin.Engine {
	router.GET(MAIN_URL, func(c *gin.Context) {
		nav := GetNav(ROOT_WIKI_PATH)

		c.HTML(
			http.StatusOK,
			"wiki.tmpl.html",
			gin.H{
				"navFilesFolders": nav,
			},
		)
	})

	router.GET(MAIN_URL+PAGE_URL+"/:id", func(c *gin.Context) {
		docId := c.Param("id")
		nav, content := GetPage(ROOT_WIKI_PATH, docId)

		c.HTML(
			http.StatusOK,
			"wiki.tmpl.html",
			gin.H{
				"navFilesFolders": nav,
				"wikiContent":     content,
			},
		)
	})

	return router
}
