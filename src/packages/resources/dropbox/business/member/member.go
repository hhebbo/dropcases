package member

func Get(limit int, includeRemoved bool) []Member {
	response := list(limit, includeRemoved)
	members := mapResponseListToMemebers(response)

	return members
}

func Add(email string, givenName string, surname string, externalId string, sendWelcomeEmail bool, Role string) Member {
	response := add(email, givenName, surname, externalId, sendWelcomeEmail, Role)
	member := mapResponseAddToMemebers(response)

	return member
}

func Merge(membersLists ...[]Member) []Member {
	var members []Member
	for _, membersSlice := range membersLists {
		members = append(members, membersSlice...)
	}

	return members
}
