package api

func SendRequestToBusinessApi(url string, dataType interface{}, authType string, arg bool) []byte {
	return send(url, dataType, authType, "", false, arg)
}

func SendRequestToEndUserApiAsUser(url string, dataType interface{}, authType string, memberTeamId string, arg bool) []byte {
	return send(url, dataType, authType, memberTeamId, false, arg)
}

func SendRequestToEndUserApiAsAdmin(url string, dataType interface{}, authType string, memberTeamId string, arg bool) []byte {
	return send(url, dataType, authType, memberTeamId, true, arg)
}
