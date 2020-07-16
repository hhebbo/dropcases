package fileFolder

import (
	"github.com/hhebbo/dropcases/src/packages/core/api"
)

const urlCreateBatch = "https://api.dropboxapi.com/2/files/create_folder_batch"

type requestParamsCreateBatch struct {
	Path       []string `json:"path"`
	AutoRename bool     `json:"autorename"`
}

func createBatch(path []string, autoRename bool, userTeamMemberId string) {
	requestParams := requestParamsCreateBatch{path, autoRename}
	api.SendRequestToEndUserApiAsUser(urlCreate, requestParams, api.AuthTypeFiles, userTeamMemberId, false)
}
