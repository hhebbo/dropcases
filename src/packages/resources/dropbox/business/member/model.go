package member

type Member struct {
	AccountId     string
	TeamMemberId  string
	Email         string
	Name          string
	Role          string
	EmailVerified bool
	Groups        []string
}

const (
	MemberRoleTeamAdmin           = "team_admin"
	MemberRoleUserManagementAdmin = "user_management_admin"
	MemberRoleSupportAdmin        = "support_admin"
	MemberRoleMemberOnly          = "member_only"
)
