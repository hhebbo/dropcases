package stateMachine

import (
	"github.com/qor/transition"
)

type Workflow struct {
	transition.Transition
	Member NewMemberRequest
}

type WorkflowState struct {
	Type     string
	Name     string
	Resource string
	Method   string
	ID       int
	Data     string
}

type WorkflowEvent struct {
	Name string
	From int
	To   int
}

type WorkflowRequest struct {
	Steps []WorkflowStep `json:"steps"`
}

type WorkflowStep struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
	Name     string `json:"name"`
	Data     string `json:"data"`
}

type NewMemberRequest struct {
	TeamMemberId string `json:"team_member_id"`
	Email string `json:"email`
	Name string `json:"name"`
}


const (
	ResourceFile           = "file"
	ResourceFileMethodCopy = "copy"

	ResourceFolder             = "folder"
	ResourceFolderMethodCreate = "create"
	ResourceFolderMethodShare  = "share"

	ResourceMember                 = "member"
	ResourceMemberMethodAdd        = "add"
	ResourceMemberMethodAddToGroup = "addToGroup"

	ResourcePaper            = "paper"
	ResourcePaperMethodShare = "share"

	StateTypeInit  = "init"
	StateTypeNode  = "node"
	StateTypeFinal = "final"
)
