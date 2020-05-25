package group

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlCreate = "https://api.dropboxapi.com/2/team/groups/create"

type requestParamsCreate struct {
	Name           string `json:"group_name"`
	ExternalId     string `json:"group_external_id"`
	ManagementType string `json:"group_management_type"`
}

type responseCreate struct {
	Id             string `json:"group_id"`
	Name           string `json:"group_name"`
	Created        string `json:"created"`
	ManagementType struct {
		Tag string `json:".tag"`
	} `json:"group_management_type"`
}

func create(name, externalId, managementType string) responseCreate {
	requestParams := requestParamsCreate{name, externalId, managementType}
	rawResponse := api.SendRequestToBusinessApi(urlCreate, requestParams, api.AuthTypeMembers, false)

	var response responseCreate
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseCreateToGroup(response responseCreate) Group {
	var group Group

	group.Id = response.Id
	group.Name = response.Name
	group.CreatedAt = response.Created
	group.ManagementType = response.ManagementType.Tag

	return group
}
