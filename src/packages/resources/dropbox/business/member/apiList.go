package member

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlList = "https://api.dropboxapi.com/2/team/members/list"

type requestParamsList struct {
	Limit          int  `json:"limit"`
	IncludeRemoved bool `json:"include_removed"`
}

type responseList struct {
	Members []responseListItem `json:"members"`
}

type responseListItem struct {
	Profile struct {
		TeamMemberId  string `json:"team_member_id"`
		AccountId     string `json:"account_id"`
		ExternalId    string `json:"external_id"`
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
		Groups         []string `json:"groups"`
		JoinedOn       string   `json:"joined_on"`
		MemberFolderID string   `json:"member_folder_id"`
	} `json:"profile"`
	Role struct {
		Tag string `json:".tag"`
	} `json:"role"`
}

func list(limit int, includeRemoved bool) responseList {
	requestParams := requestParamsList{limit, includeRemoved}
	rawResponse := api.SendRequestToBusinessApi(urlList, requestParams, api.AuthTypeMembers, false)

	var response responseList
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseListToMemebers(response responseList) []Member {
	var members []Member

	for _, rawMember := range response.Members {
		var member Member
		member.AccountId = rawMember.Profile.AccountId
		member.TeamMemberId = rawMember.Profile.TeamMemberId
		member.Email = rawMember.Profile.Email
		member.Name = rawMember.Profile.Name.DisplayName
		member.Role = rawMember.Role.Tag
		member.EmailVerified = rawMember.Profile.EmailVerified
		member.Groups = rawMember.Profile.Groups

		members = append(members, member)
	}

	return members
}
