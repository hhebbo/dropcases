package group

type Group struct {
	Id             string
	Name           string
	CreatedAt      string
	ManagementType string
}

const (
	MemberAccessTypeMember = "member"
	MemberAccessTypeOwner  = "owner"
)

type GroupMember struct {
	TeamMemberId string
	AccessType   string
}
