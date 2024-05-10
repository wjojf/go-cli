package actions

import "github.com/wjojf/go-ssh-tui/internal/types"

type FilterableAction struct {
	Action types.Action
}

func (a FilterableAction) FilterValue() string {
	return a.Action.Name()
}

func (a FilterableAction) Title() string {
	return a.Action.Name()
}

func (a FilterableAction) Description() string {
	return a.Action.Description()
}

func FromActionList(l types.ActionList) []FilterableAction {
	var actions []FilterableAction
	for _, a := range l {
		actions = append(actions, FilterableAction{Action: a})
	}
	return actions
}
