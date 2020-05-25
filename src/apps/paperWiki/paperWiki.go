package paperWiki

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/src/packages/resources/dropbox/endUser/fileFolder"
)

func GetRoutes(router *gin.Engine) *gin.Engine {
	return getRoutes(router)
}

func GetNav(path string) []fileFolder.FileFolder {
	sections, pages := getNavSectionsAndPages(path)
	nav := sortSectionsAndPages(sections, pages)

	return nav
}

func GetPage(path string, docId string) ([]fileFolder.FileFolder, string) {
	nav := GetNav(path)
	pageContent := getPageContent(docId)

	return nav, pageContent
}
