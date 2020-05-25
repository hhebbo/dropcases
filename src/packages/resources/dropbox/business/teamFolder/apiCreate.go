package teamFolder

import (
	"encoding/json"

	"github.com/hhebbo/dropcases/src/packages/core/api"
)

const urlCreate = "https://api.dropboxapi.com/2/team/team_folder/create"

type requestParamsCreate struct {
	Name string `json:"name"`
}

type responseCreate struct {
	TeamFolderID string `json:"team_folder_id"`
	Name         string `json:"name"`
	Status       struct {
		Tag string `json:".tag"`
	} `json:"status"`
	IsTeamSharedDropbox string `json:"is_team_shared_dropbox"`
	SyncSetting         struct {
		Tag string `json:".tag"`
	} `json:"sync_setting"`
	ContentSyncSettings []string `json:"content_sync_settings"`
}

func create(name string) responseCreate {
	requestParams := requestParamsCreate{name}
	rawResponse := api.SendRequestToBusinessApi(urlCreate, requestParams, api.AuthTypeFiles, false)

	var response responseCreate
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseCreateToTeamFolder(response responseCreate) TeamFolder {
	var teamFolder TeamFolder

	teamFolder.Id = response.TeamFolderID
	teamFolder.Name = response.Name
	teamFolder.IsTeamSharedDropbox = response.IsTeamSharedDropbox

	return teamFolder
}
