package paperWiki

import (
	"strings"

	"github.com/hhebbo/dropcases/src/apps/appConfig"
	"github.com/hhebbo/dropcases/src/packages/core/config"
	"github.com/hhebbo/dropcases/src/packages/resources/dropbox/endUser/fileFolder"
)

func getNavSectionsAndPages(path string) ([]fileFolder.FileFolder, []fileFolder.FileFolder) {
	var sections []fileFolder.FileFolder
	var pages []fileFolder.FileFolder

	adminMemberId := config.GetValue(appConfig.DROPCASES_ADMIN_MEMBER_ID)
	filesFolders := fileFolder.List(path, adminMemberId)

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
