package stateMachine

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

func build(states []WorkflowState, events []WorkflowEvent) *transition.StateMachine {
	sm := transition.New(&Workflow{})
	sm = buildStates(sm, states)
	sm = buildEvents(sm, events, states)

	return sm
}

func buildStates(workflow *transition.StateMachine, states []WorkflowState) *transition.StateMachine {
	for _, s := range states {
		state := s
		if state.Type == StateTypeInit {
			workflow.Initial(state.Name)
		} else {
			workflow.State(state.Name).Enter(func(r interface{}, tx *gorm.DB) error {
				runResourceMethod(state.Resource, state.Method, r.(*Workflow))
				return nil
			})
		}
	}

	return workflow
}

func buildEvents(workflow *transition.StateMachine, events []WorkflowEvent, states []WorkflowState) *transition.StateMachine {
	for _, event := range events {
		workflow.Event(event.Name).To(states[event.To].Name).From(states[event.From].Name)
	}

	return workflow
}
