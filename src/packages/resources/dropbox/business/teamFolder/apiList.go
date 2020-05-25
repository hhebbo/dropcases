package teamFolder

import (
	"encoding/json"

	"github.com/hhebbo/dropcases/src/packages/core/api"
)

const urlList = "https://api.dropboxapi.com/2/team/team_folder/list"

type requestParamsList struct {
	Limit int `json:"limit"`
}

type responseList struct {
	TeamFolders []TeamFolder `json:"team_folders"`
}

func list(limit int) []TeamFolder {
	requestParams := requestParamsList{limit}
	rawResponse := api.SendRequestToBusinessApi(urlList, requestParams, api.AuthTypeFiles, false)

	var response responseList
	json.Unmarshal(rawResponse, &response)

	return response.TeamFolders
}
