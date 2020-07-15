package templates

import (
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
	projectName := "MediaPro Video Production"
	template := "Workflow1"
	adminMemberId := config.GetValue(appConfig.DROPCASES_ADMIN_MEMBER_ID)
	groupId := "g:65b86a75b6caefa00000000000000dc5"

	tf := teamFolder.Create(projectName)
	var gs []teamFolder.TeamFolderGroup
	g := teamFolder.TeamFolderGroup{groupId, teamFolder.GroupAccessLevelEditor}
	gs = append(gs, g)
	teamFolder.AddGroups(tf.Id, gs, false, adminMemberId)

	folders, _ := fileFolder.List("/FolderTemplates/"+template, adminMemberId)
	for _, f := range folders[1:] {
		var pathSegments = strings.Split(f.Path, "/")
		var path = "/" + projectName + "/" + strings.Title(strings.Join(pathSegments[3:], "/"))
		fileFolder.CreateFolder(path, false, adminMemberId)
	}
}
