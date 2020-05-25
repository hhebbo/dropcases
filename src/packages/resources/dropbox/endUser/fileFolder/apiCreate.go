package fileFolder

import (
	"encoding/json"

	"github.com/hhebbo/dropcases/src/packages/core/api"
)

const urlCreate = "https://api.dropboxapi.com/2/files/create_folder_v2"

type requestParamsCreate struct {
	Path       string `json:"path"`
	AutoRename bool   `json:"autorename"`
}

type responseCreate struct {
	Metadata struct {
		Name        string `json:"name"`
		Id          string `json:"id"`
		PathLower   string `json:"path_lower"`
		PathDisplay string `json:"path_display"`
		SharingInfo struct {
			ReadOnly             bool   `json:"read_only"`
			SharedFolderId       string `json:"shared_folder_id"`
			ParentSharedFolderId string `json:"parent_shared_folder_id"`
			TraverseOnly         bool   `json:"traverse_only"`
			NoAccess             bool   `json:"no_access"`
		} `json:"sharing_info"`
		PropertyGroups []responseFieldsItem `json:"property_groups"`
	} `json:"metadata"`
}

type responsePropertyGroupsItem struct {
	TemplateId string               `json:"template_id"`
	Fields     []responseFieldsItem `json:"fields"`
}

type responseFieldsItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func create(path string, autoRename bool, userTeamMemberId string) responseCreate {
	requestParams := requestParamsCreate{path, autoRename}
	rawResponse := api.SendRequestToEndUserApiAsUser(urlCreate, requestParams, api.AuthTypeFiles, userTeamMemberId, false)

	var response responseCreate
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseCreateToFileFolder(response responseCreate) FileFolder {
	var fileFolder FileFolder

	fileFolder.Id = response.Metadata.Id
	fileFolder.Name = response.Metadata.Name
	fileFolder.Path = response.Metadata.PathDisplay
	fileFolder.SharedFolderId = response.Metadata.SharingInfo.SharedFolderId
	fileFolder.ParentFolderId = response.Metadata.SharingInfo.ParentSharedFolderId

	return fileFolder
}
