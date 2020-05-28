package fileFolder

func CreateFolder(path string, autoRename bool, userTeamMemberId string) FileFolder {
	response := create(path, autoRename, userTeamMemberId)
	folder := mapResponseCreateToFileFolder(response)

	return folder
}

func ShareFolder(path string, inviteeTeamMemberId string, accessLevel string, userTeamMemberId string) {
	share(path, inviteeTeamMemberId, accessLevel, userTeamMemberId)
}

func Copy(from string, to string, allowSharedFolder bool, autorename bool, allowOwnershipTransfer bool, userTeamMemberId string) FileFolder {
	response := copy(from, to, allowSharedFolder, autorename, allowOwnershipTransfer, userTeamMemberId)
	fileFolder := mapResponseCopyToFileFolder(response)

	return fileFolder
}

func List(path string, adminMemberTeamId string) ([]FileFolder, string) {
	return list(path, adminMemberTeamId)
}

func ListContinue(cursor string, adminMemberTeamId string) []FileFolder {
	return listContinue(cursor, adminMemberTeamId)
}

func GetFileContent(id string, userMemberTeamId string) string {
	return getFileContent(id, userMemberTeamId)
}
