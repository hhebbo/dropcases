package teamFolder

type TeamFolder struct {
	Id                  string
	Name                string
	IsTeamSharedDropbox string
}

const (
	GroupAccessLevelOwner           = "owner"
	GroupAccessLevelEditor          = "editor"
	GroupAccessLevelViewer          = "viewer"
	GroupAccessLevelViewerNoComment = "viewer_no_comment"
)

type TeamFolderGroup struct {
	Id          string
	AccessLevel string
}
