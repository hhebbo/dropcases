package stateMachine

import (
	"fmt"

	"github.com/qor/transition"
)

var sm *transition.StateMachine
var events []WorkflowEvent

func Build(wReq WorkflowRequest) {
	// fmt.Println(wReq)

	var statesDynamic []WorkflowState
	var eventsDynamic []WorkflowEvent

	for i, step := range wReq.Steps {
		stepType := StateTypeNode
		if i == 0 {
			stepType = StateTypeInit
		} else if i+1 == len(wReq.Steps) {
			stepType = StateTypeFinal
		}

		statesDynamic = append(statesDynamic, WorkflowState{stepType, step.Name, step.Resource, step.Action, i, step.Data})

		if i > 0 {
			eventsDynamic = append(eventsDynamic, WorkflowEvent{step.Name, i - 1, i})
		}
	}

	sm = build(statesDynamic, eventsDynamic)
	events = eventsDynamic
}

func HandleNewMember(newMember NewMemberRequest) {
	fmt.Println("          NEW MEMBER TRIGGER          ")
	fmt.Println(newMember)

	wf := &Workflow{Member: newMember}
	go internalRun(wf)
}

func Run() {
	wf := &Workflow{}
	internalRun(wf)
}

func internalRun(wf *Workflow) {
	for _, event := range events {
		fmt.Println("                  |                  ")
		fmt.Println("                  |                  ")

		sm.Trigger(event.Name, wf, nil)
		fmt.Println("          TRIGGERING EVENT          ")
		fmt.Println(event.Name)

		fmt.Println("                  |                  ")
		fmt.Println("                  |                  ")
		fmt.Println("                  V                  ")

		fmt.Println("-------------------------------------")
		fmt.Println("|           CURRENT STATE           |")
		fmt.Println("| " + wf.GetState() + " |")
		fmt.Println("-------------------------------------")
	}

	fmt.Println("-------------------------")
	fmt.Println("|         DONE!         |")
	fmt.Println("|      Final state      |")
	fmt.Println("-------------------------")
}

func BuildPreDefinedDemo() {
	events = getEvents()
	sm = build(getStates(), events)
}

func getStates() []WorkflowState {
	var states []WorkflowState
	states = append(states, WorkflowState{StateTypeInit, "New member confirmed", "", "", 0, ""})
	states = append(states, WorkflowState{StateTypeNode, "Member added to team", ResourceMember, ResourceMemberMethodAdd, 1, ""})
	states = append(states, WorkflowState{StateTypeNode, "Member added to Engineering group", ResourceMember, ResourceMemberMethodAddToGroup, 2, ""})
	states = append(states, WorkflowState{StateTypeNode, "Folder created", ResourceFolder, ResourceFolderMethodCreate, 3, ""})
	states = append(states, WorkflowState{StateTypeNode, "Files moved into folder", ResourceFile, ResourceFileMethodCopy, 4, ""})
	states = append(states, WorkflowState{StateTypeNode, "Paper shared", ResourcePaper, ResourcePaperMethodShare, 5, ""})
	states = append(states, WorkflowState{StateTypeFinal, "Folder shared", ResourceFolder, ResourceFolderMethodShare, 6, ""})

	return states
}

func getEvents() []WorkflowEvent {
	var events []WorkflowEvent
	events = append(events, WorkflowEvent{"Add member to team", 0, 1})
	events = append(events, WorkflowEvent{"Add member to Engineering group", 1, 2})
	events = append(events, WorkflowEvent{"Create folder in Members team folder", 2, 3})
	events = append(events, WorkflowEvent{"Copy onboarding files fo folder", 3, 4})
	events = append(events, WorkflowEvent{"Share onboarding paper doc", 4, 5})
	events = append(events, WorkflowEvent{"Share folder with member", 5, 6})

	return events
}
