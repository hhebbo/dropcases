package fileFolder

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlList = "https://api.dropboxapi.com/2/files/list_folder"

type requestParamsList struct {
	Path                            string `json:"path"`
	Recursive                       bool   `json:"recursive"`
	IncludeMediaInfo                bool   `json:"include_media_info"`
	IncludeDeleted                  bool   `json:"include_deleted"`
	IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members"`
	IncludeMountedFolders           bool   `json:"include_mounted_folders"`
	IncludeNonDownloadableFiles     bool   `json:"include_non_downloadable_files"`
}

type responseList struct {
	FileFolders []FileFolder `json:"entries"`
}

func list(path string, adminMemberTeamId string) []FileFolder {
	requestParams := requestParamsList{path, true, true, false, true, true, true}
	rawResponse := api.SendRequestToEndUserApiAsAdmin(urlList, requestParams, api.AuthTypeFiles, adminMemberTeamId, false)

	var response responseList
	json.Unmarshal(rawResponse, &response)

	return response.FileFolders
}
