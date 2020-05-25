package fileFolder

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlCopy = "https://api.dropboxapi.com/2/files/copy_v2"

type requestParamsCopy struct {
	FromPath               string `json:"from_path"`
	ToPath                 string `json:"to_path"`
	AllowSharedFolder      bool   `json:"allow_shared_folder"`
	AutoRename             bool   `json:"autorename"`
	AllowOwnershipTransfer bool   `json:"allow_ownership_transfer"`
}

type responseCopy struct {
	Metadata struct {
		Tag            string `json:".tag"`
		Name           string `json:"name"`
		Id             string `json:"id"`
		ClientModified string `json:"client_modified"`
		ServerModified string `json:"server_modified"`
		Rev            string `json:"rev"`
		Size           int    `json:"size"`
		PathLower      string `json:"path_lower"`
		PathDisplay    string `json:"path_display"`
		SharingInfo    struct {
			ReadOnly             bool   `json:"read_only"`
			ParentSharedFolderId string `json:"parent_shared_folder_id"`
			ModifiedBy           string `json:"modified_by"`
		} `json:"sharing_info"`
		IsDownloadable           bool                 `json:"is_downloadable"`
		PropertyGroups           []responseFieldsItem `json:"property_groups"`
		HasExplicitSharedMembers bool                 `json:"has_explicit_shared_members"`
		ContentHash              string               `json:"content_hash"`
	} `json:"metadata"`
}

func copy(from string, to string, allowSharedFolder bool, autorename bool, allowOwnershipTransfer bool, userTeamMemberId string) responseCopy {
	requestParams := requestParamsCopy{from, to, allowSharedFolder, autorename, allowOwnershipTransfer}
	rawResponse := api.SendRequestToEndUserApiAsUser(urlCopy, requestParams, api.AuthTypeFiles, userTeamMemberId, false)

	var response responseCopy
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseCopyToFileFolder(response responseCopy) FileFolder {
	var fileFolder FileFolder

	fileFolder.Id = response.Metadata.Id
	fileFolder.Name = response.Metadata.Name
	fileFolder.Path = response.Metadata.PathDisplay
	fileFolder.ParentFolderId = response.Metadata.SharingInfo.ParentSharedFolderId

	return fileFolder
}
