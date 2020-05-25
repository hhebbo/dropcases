package fileFolder

import (
	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlFileContent = "https://content.dropboxapi.com/2/files/export"

type requestParamsFileContent struct {
	Path string `json:"path"`
}

type responseFileContent struct {
	content string
}

func getFileContent(id string, userMemberTeamId string) string {
	requestParams := requestParamsFileContent{id}
	rawResponse := api.SendRequestToEndUserApiAsUser(urlFileContent, requestParams, api.AuthTypeFiles, userMemberTeamId, true)

	return string(rawResponse)
}
