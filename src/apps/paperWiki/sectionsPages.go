package paperWiki

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hhebbo/dropcases/src/apps/appConfig"
	"github.com/hhebbo/dropcases/src/packages/core/config"
	"github.com/hhebbo/dropcases/src/packages/core/fileCache"
	"github.com/hhebbo/dropcases/src/packages/resources/dropbox/endUser/fileFolder"
)

const (
	FILECACHE_PATH = "paperWiki/folderListCursor.txt"
)

func getNavSectionsAndPages(path string) ([]fileFolder.FileFolder, []fileFolder.FileFolder) {
	var sections []fileFolder.FileFolder
	var pages []fileFolder.FileFolder

	filesFolders := getFilesFolders(path)

	i := 1
	for _, f := range filesFolders[1:] {
		if f.Type == fileFolder.FileFolderTypeFolder {
			sections = append(sections, filesFolders[i])
		}

		if f.Type == fileFolder.FileFolderTypeFile {
			filesFolders[i].Name = strings.Split(f.Name, ".")[0]
			filesFolders[i].Url = MAIN_URL + PAGE_URL + "/" + f.Id
			filesFolders[i].Parent = strings.Split(f.Path, "/")[2]

			pages = append(pages, filesFolders[i])
		}
		i++
	}

	return sections, pages
}

func sortSectionsAndPages(sections []fileFolder.FileFolder, pages []fileFolder.FileFolder) []fileFolder.FileFolder {
	var nav []fileFolder.FileFolder

	for _, section := range sections {
		nav = append(nav, section)
		for _, page := range pages {
			if strings.ToLower(section.Name) == page.Parent {
				nav = append(nav, page)
			}
		}
	}

	return nav
}

func getPageContent(docId string) string {
	adminMemberId := config.GetValue(appConfig.DROPCASES_ADMIN_MEMBER_ID)
	content := fileFolder.GetFileContent(docId, adminMemberId)

	return content
}

func getFilesFolders(path string) []fileFolder.FileFolder {
	adminMemberId := config.GetValue(appConfig.DROPCASES_ADMIN_MEMBER_ID)

	fmt.Println(strconv.FormatBool(fileCache.Exists(FILECACHE_PATH)))

	var filesFolders []fileFolder.FileFolder
	var cursor string
	if !fileCache.Exists(FILECACHE_PATH) {
		filesFolders, cursor = fileFolder.List(path, adminMemberId)
		fileCache.Save(FILECACHE_PATH, cursor)
	} else {
		cursor = fileCache.Get(FILECACHE_PATH)
	}
	filesFolders = fileFolder.ListContinue(cursor, adminMemberId)

	return filesFolders
}
