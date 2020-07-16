package templates

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hhebbo/dropcases/src/apps/appConfig"
	"github.com/hhebbo/dropcases/src/packages/core/config"
	"github.com/hhebbo/dropcases/src/packages/resources/dropbox/business/teamFolder"
	"github.com/hhebbo/dropcases/src/packages/resources/dropbox/endUser/fileFolder"
)

func GetRoutes(router *gin.Engine) *gin.Engine {
	return getRoutes(router)
}

func CreateFoldersFromTemplate() {
	projectName := "4"
	template := "internal"
	adminMemberId := config.GetValue(appConfig.DROPCASES_ADMIN_MEMBER_ID)
	group1Id := "g:65b86a75b6caefa000000000000001e7"

	tf := teamFolder.Create(projectName)
	var gs []teamFolder.TeamFolderGroup
	g1 := teamFolder.TeamFolderGroup{group1Id, teamFolder.GroupAccessLevelEditor}
	gs = append(gs, g1)
	teamFolder.AddGroups(tf.Id, gs, false, adminMemberId)

	fromPath := "/Templates/" + template
	folders, _ := fileFolder.List(fromPath, adminMemberId)

	for _, f := range folders[1:] {
		var pathSegments = strings.Split(f.Path, "/")
		var toPath = "/" + projectName + "/"

		fmt.Println(f.Name)
		if f.Type == fileFolder.FileFolderTypeFolder {
			folderPath := toPath + strings.Title(strings.Join(pathSegments[3:], "/"))
			fileFolder.CreateFolder(folderPath, false, adminMemberId)
		}

		if f.Type == fileFolder.FileFolderTypeFile {
			filePath := toPath + f.Name
			fileFolder.Copy(fromPath+"/"+f.Name, filePath, true, false, true, adminMemberId)
		}

	}
}
