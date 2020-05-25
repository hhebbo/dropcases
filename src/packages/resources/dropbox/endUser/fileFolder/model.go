package fileFolder

type FileFolder struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path_lower"`
	SharedFolderId string `json:"path_display"`
	ParentFolderId string `json:"parent_shared_id"`
	Type           string `json:".tag"`
	Content        string
	Url            string
	Parent         string
}

const (
	FileFolderTypeFolder             = "folder"
	FileFolderTypeFile               = "file"
	MemberAccessLevelOwner           = "owner"
	MemberAccessLevelEditor          = "editor"
	MemberAccessLevelViewer          = "viewer"
	MemberAccessLevelViewerNoComment = "viewer_no_comment"
)

type FolderExternalMember struct {
	Email       string
	AccessLevel string
}
