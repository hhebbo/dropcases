package folderTemplates

import (
	"strings"

	"github.com/heroku/go-getting-started/src/packages/resources/dropbox/business/teamFolder"
	filefolder "github.com/heroku/go-getting-started/src/packages/resources/dropbox/endUser/fileFolder"
)

func CreateFoldersFromTemplate() {
	projectName := "MediaPro Video Production 1"
	template := "Project"
	adminTeamMemberId := "dbmid:AAC-CvicHJKg7ND_MGrDqgxm1rBa0lmWV28"
	groupId := "g:65b86a75b6caefa000000000000001e7"

	tf := teamFolder.Create(projectName)
	var gs []teamFolder.TeamFolderGroup
	g := teamFolder.TeamFolderGroup{groupId, teamFolder.GroupAccessLevelEditor}
	gs = append(gs, g)
	teamFolder.AddGroups(tf.Id, gs, false, adminTeamMemberId)

	folders := filefolder.List("/FolderTemplates/"+template, adminTeamMemberId)
	for _, f := range folders[1:] {
		var pathSegments = strings.Split(f.Path, "/")
		var path = "/" + projectName + "/" + strings.Join(pathSegments[3:], "/")
		filefolder.CreateFolder(path, false, adminTeamMemberId)
	}
}
