package teamFolder

func Create(name string) TeamFolder {
	createResponse := create(name)
	teamFolder := mapResponseCreateToTeamFolder(createResponse)

	return teamFolder
}

func AddGroups(teamFolderId string, groups []TeamFolderGroup, quite bool, adminTeamMemberId string) {
	addGroups(teamFolderId, groups, quite, adminTeamMemberId)
}

func List(limit int) []TeamFolder {
	return list(limit)
}

func Count() int {
	return len(List(1000))
}
