package fileFolder

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlShare = "https://api.dropboxapi.com/2/sharing/share_folder"
const urlAddMember = "/sharing/add_folder_member"

type requestParamsShare struct {
	Path string `json:"path"`
}

type responseShare struct {
	ShareFolderId string `json:"shared_folder_id"`
}

type requestParamsAddMember struct {
	ShareFolderId string                    `json:"shared_folder_id"`
	Members       []requestParamsMemberItem `json:"members"`
	Quiet         bool                      `json:"quiet"`
}

type requestParamsMemberItem struct {
	Member struct {
		Tag       string `json:".tag"`
		DropboxId string `json:"dropbox_id"`
	} `json:"member"`
	AccessLevel string `json:"access_level"`
}

func share(path string, inviteeTeamMemberId string, accessLevel string, userTeamMemberId string) {
	requestParamsShare := requestParamsShare{path}
	rawResponseShare := api.SendRequestToEndUserApiAsUser(urlShare, requestParamsShare, api.AuthTypeFiles, userTeamMemberId, false)

	var responseShare responseShare
	json.Unmarshal(rawResponseShare, &responseShare)

	var requestParamsMemberItem requestParamsMemberItem
	requestParamsMemberItem.Member.Tag = "dropbox_id"
	requestParamsMemberItem.Member.DropboxId = inviteeTeamMemberId
	requestParamsMemberItem.AccessLevel = accessLevel

	var requestParamsAddMember requestParamsAddMember
	requestParamsAddMember.ShareFolderId = responseShare.ShareFolderId
	requestParamsAddMember.Members = append(requestParamsAddMember.Members, requestParamsMemberItem)
	requestParamsAddMember.Quiet = true

	api.SendRequestToEndUserApiAsUser(urlAddMember, requestParamsAddMember, api.AuthTypeFiles, userTeamMemberId, false)
}
