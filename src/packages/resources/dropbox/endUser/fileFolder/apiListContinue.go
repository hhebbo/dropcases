package fileFolder

import (
	"encoding/json"

	"github.com/hhebbo/dropcases/src/packages/core/api"
)

const urlListContinue = "https://api.dropboxapi.com/2/files/list_folder/continue"

type requestParamsListContinue struct {
	Cursor string `json:"cursor"`
}

type responseListContinue struct {
	FileFolders []FileFolder `json:"entries"`
}

func listContinue(cursor string, adminMemberTeamId string) []FileFolder {
	requestParams := requestParamsListContinue{cursor}
	rawResponse := api.SendRequestToEndUserApiAsAdmin(urlListContinue, requestParams, api.AuthTypeFiles, adminMemberTeamId, false)

	var response responseListContinue
	json.Unmarshal(rawResponse, &response)

	return response.FileFolders
}
