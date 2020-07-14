package stateMachine

import (
	"time"

	"github.com/heroku/go-getting-started/src/packages/resources/dropbox/business/group"
	"github.com/heroku/go-getting-started/src/packages/resources/dropbox/endUser/fileFolder"
	"github.com/heroku/go-getting-started/src/packages/resources/dropbox/endUser/paper"
)

const (
	TeamAdminId       = "dbmid:AAFfCiGwd877ecDoanNO5kwYlK14hPYEvCY"
	GroupId           = "g:5f9ff9897c43df400000000000000036"
	PaperDocId        = "PAKiYxyEPDdF8AkeGpFpz"
	SourceFolder      = "/Onboarding/"
	DestinationFolder = "/Members/"
)

func runResourceMethod(resource string, method string, data *Workflow) {
	switch resource {
	case ResourceFile:
		handleFileResource(method, data)

	case ResourceFolder:
		handleFolderResource(method, data)

	case ResourceMember:
		handleMemberResource(method, data)

	case ResourcePaper:
		handlePaperResource(method, data)
	}
}

func handleFileResource(method string, data *Workflow) {
	switch method {
	case ResourceFileMethodCopy:
		files := fileFolder.List(SourceFolder, TeamAdminId)
		for _, f := range files {
			fileFolder.Copy(SourceFolder+f.Name, DestinationFolder+data.Member.Name+"/"+f.Name, false, false, false, TeamAdminId)
		}
	}
}

func handleFolderResource(method string, data *Workflow) {
	switch method {
	case ResourceFolderMethodCreate:
		fileFolder.CreateFolder(DestinationFolder+data.Member.Name, true, TeamAdminId)
		time.Sleep(7 * time.Second)

	case ResourceFolderMethodShare:
		fileFolder.ShareFolder(DestinationFolder+data.Member.Name, data.Member.TeamMemberId, fileFolder.MemberAccessLevelEditor, TeamAdminId)
	}
}

func handleMemberResource(method string, data *Workflow) {
	switch method {
	case ResourceMemberMethodAdd:
		// member.Add("hussam+testuser7@hanfordinc.com", "H", "H", "", false, member.MemberRoleMemberOnly)
		//var g []string
		//data.Member = member.Member{"", MemberTeamMemberID, MemberEmail, MemberName, "", true, g}

	case ResourceMemberMethodAddToGroup:
		var groupMembers []group.GroupMember
		groupMember := group.GroupMember{data.Member.TeamMemberId, group.MemberAccessTypeMember}
		groupMembers = append(groupMembers, groupMember)
		group.AddMembers(GroupId, groupMembers, false)
	}
}

func handlePaperResource(method string, data *Workflow) {
	switch method {
	case ResourcePaperMethodShare:
		paper.Share(PaperDocId, data.Member.TeamMemberId, paper.MemberPermissionLevelViewAndComment, TeamAdminId)
	}
}
