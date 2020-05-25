package group

func Create(name, externalId, managementType string) Group {
	response := create(name, externalId, managementType)
	group := mapResponseCreateToGroup(response)

	return group
}

func AddMembers(groupId string, members []GroupMember, returnMembers bool) Group {
	response := addMembers(groupId, members, returnMembers)
	group := mapResponseAddMembersToGroup(response)

	return group
}
