package group

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlAddMembers = "https://api.dropboxapi.com/2/team/groups/members/add"

type requestParamsAddMembers struct {
	Group struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	} `json:"group"`
	Members       []requestParamsMember `json:"members"`
	ReturnMembers bool                  `json:"return_members"`
}

type requestParamsMember struct {
	User struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	} `json:"user"`
	AccessType string `json:"access_type"`
}

type responseAddMembers struct {
	GroupInfo struct {
		GroupName      string `json:"group_name"`
		GroupId        string `json:"group_id"`
		Created        string `json:"created"`
		ManagementType struct {
			Tag string `json:".tag"`
		} `json:"group_management_type"`
		MemberCount int `json:"member_count"`
		Members     []responseMember
	} `json:"group_info"`
	AsyncJobId string `json:"async_job_id"`
}

type responseMember struct {
	Profile struct {
		TeamMemberId  string `json:"team_member_id"`
		AccountId     string `json:"account_id"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Status        struct {
			Tag string `json:".tag"`
		} `json:"status"`
		Name struct {
			GivenName       string `json:"given_name"`
			Surname         string `json:"surname"`
			FamiliarName    string `json:"familiar_name"`
			DisplayName     string `json:"display_name"`
			AbbreviatedName string `json:"abbreviated_name"`
		} `json:"name"`
		MembershipType struct {
			Tag string `json:".tag"`
		} `json:"membership_type"`
		JoinedOn    string `json:"joined_on"`
		AccessTuype struct {
			Tag string `json:".tag"`
		} `json:"access_type"`
	} `json:"profile"`
}

func addMembers(groupId string, members []GroupMember, retrunMembers bool) responseAddMembers {
	requestParams := requestParamsAddMembers{}
	requestParams.Group.Tag = "group_id"
	requestParams.Group.GroupId = groupId
	for _, groupMember := range members {
		var requestParamsMember requestParamsMember
		requestParamsMember.User.Tag = "team_member_id"
		requestParamsMember.User.TeamMemberId = groupMember.TeamMemberId
		requestParamsMember.AccessType = groupMember.AccessType

		requestParams.Members = append(requestParams.Members, requestParamsMember)
	}
	requestParams.ReturnMembers = retrunMembers

	rawResponse := api.SendRequestToBusinessApi(urlAddMembers, requestParams, api.AuthTypeMembers, false)

	var response responseAddMembers
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseAddMembersToGroup(response responseAddMembers) Group {
	var group Group

	group.Id = response.GroupInfo.GroupId
	group.Name = response.GroupInfo.GroupName
	group.CreatedAt = response.GroupInfo.Created
	group.ManagementType = response.GroupInfo.ManagementType.Tag

	return group
}
