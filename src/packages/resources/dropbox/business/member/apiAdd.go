package member

import (
	"encoding/json"

	"github.com/heroku/go-getting-started/src/packages/core/api"
)

const urlAdd = "https://api.dropboxapi.com/2/team/members/add"

type requestParamsAdd struct {
	NewMembers []requestParamsAddItem `json:"new_members"`
}

type requestParamsAddItem struct {
	Email            string `json:"member_email"`
	GivenName        string `json:"member_given_name"`
	Surname          string `json:"member_surname"`
	ExternalID       string `json:"member_external_id"`
	SendWelcomeEmail bool   `json:"send_welcome_email"`
	Role             string `json:"role"`
}

type responseAdd struct {
	Tag      string             `json:".tag"`
	Complete []responseListItem `json:"complete"`
}

type responseAddItem struct {
	Profile struct {
		Tag           string `json:".tag"`
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

func add(email string, givenName string, surname string, externalId string, sendWelcomeEmail bool, Role string) responseAdd {
	newMember := requestParamsAddItem{email, givenName, surname, externalId, sendWelcomeEmail, Role}
	requestParams := requestParamsAdd{}
	requestParams.NewMembers = append(requestParams.NewMembers, newMember)
	rawResponse := api.SendRequestToBusinessApi(urlAdd, requestParams, api.AuthTypeMembers, false)

	var response responseAdd
	json.Unmarshal(rawResponse, &response)

	return response
}

func mapResponseAddToMemebers(response responseAdd) Member {
	var member Member

	member.AccountId = response.Complete[0].Profile.AccountId
	member.TeamMemberId = response.Complete[0].Profile.TeamMemberId
	member.Email = response.Complete[0].Profile.Email
	member.Name = response.Complete[0].Profile.Name.DisplayName
	member.Role = response.Complete[0].Role.Tag
	member.EmailVerified = response.Complete[0].Profile.EmailVerified
	member.Groups = response.Complete[0].Profile.Groups

	return member
}
