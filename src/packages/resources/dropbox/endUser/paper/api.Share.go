package paper

import (
	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlAddMember = "https://api.dropboxapi.com/2/paper/docs/users/add"

type requestParamsAddMember struct {
	DocId   string                    `json:"doc_id"`
	Members []requestParamsMemberItem `json:"members"`
	Quiet   bool                      `json:"quiet"`
}

type requestParamsMemberItem struct {
	Member struct {
		Tag       string `json:".tag"`
		DropboxId string `json:"dropbox_id"`
	} `json:"member"`
	PermissionLevel string `json:"permission_level"`
}

func share(docId string, inviteeTeamMemberId string, permissionLevel string, userTeamMemberId string) {
	var requestParamsMemberItem requestParamsMemberItem
	requestParamsMemberItem.Member.Tag = "dropbox_id"
	requestParamsMemberItem.Member.DropboxId = inviteeTeamMemberId
	requestParamsMemberItem.PermissionLevel = permissionLevel

	var requestParamsAddMember requestParamsAddMember
	requestParamsAddMember.DocId = docId
	requestParamsAddMember.Members = append(requestParamsAddMember.Members, requestParamsMemberItem)
	requestParamsAddMember.Quiet = true

	api.SendRequestToEndUserApiAsUser(urlAddMember, requestParamsAddMember, api.AuthTypeFiles, userTeamMemberId, false)
}
