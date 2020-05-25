package teamFolder

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlAddGroup = "https://api.dropboxapi.com/2/sharing/add_folder_member"

type requestParamsAddGroup struct {
	TeamFolderId  string               `json:"shared_folder_id"`
	Groups        []requestParamsGroup `json:"members"`
	Quite         bool                 `json:"quiet"`
	CustomMessage string               `json:"custom_message"`
}

type requestParamsGroup struct {
	Group struct {
		Tag     string `json:".tag"`
		GroupId string `json:"dropbox_id"`
	} `json:"member"`
	AccessLevel string `json:"access_level"`
}

func addGroups(teamFolderId string, groups []TeamFolderGroup, quite bool, adminTeamMemberId string) {
	requestParams := requestParamsAddGroup{}
	requestParams.TeamFolderId = teamFolderId
	for _, group := range groups {
		var requestParamsGroup requestParamsGroup
		requestParamsGroup.Group.Tag = "dropbox_id"
		requestParamsGroup.Group.GroupId = group.Id
		requestParamsGroup.AccessLevel = group.AccessLevel

		requestParams.Groups = append(requestParams.Groups, requestParamsGroup)
	}
	requestParams.Quite = quite
	requestParams.CustomMessage = "New groups added to a team folder"

	rawResponse := api.SendRequestToEndUserApiAsAdmin(urlAddGroup, requestParams, api.AuthTypeFiles, adminTeamMemberId, false)

	var response responseCreate
	json.Unmarshal(rawResponse, &response)
}
