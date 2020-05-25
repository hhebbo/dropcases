package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	requestTypePost = "POST"
	AuthTypeMembers = "MEMBERS"
	AuthTypeFiles   = "FILES"
)

func send(url string, requestParams interface{}, authType string, memberTeamId string, needsAdmin bool, arg bool) []byte {
	apiToken := getToken(authType)
	payloadByte, _ := json.Marshal(requestParams)
	payloadBuffer := bytes.NewBuffer(payloadByte)

	var request *http.Request
	if arg {
		request, _ = http.NewRequest(requestTypePost, url, nil)
		request.Header.Add("Dropbox-API-Arg", string(payloadByte))
	} else {
		request, _ = http.NewRequest(requestTypePost, url, payloadBuffer)
		request.Header.Add("Content-Type", "application/json")
	}

	request.Header.Add("Authorization", "Bearer "+apiToken)

	if memberTeamId != "" {
		if needsAdmin {
			request.Header.Add("Dropbox-API-Select-Admin", memberTeamId)
		} else {
			request.Header.Add("Dropbox-API-Select-User", memberTeamId)
		}
	}

	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return body
}
