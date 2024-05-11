package actions

import (
	"github.com/wjojf/go-ssh-tui/internal/types/action"
)

type ItemAction struct {
	Action action.Action
}

func (a ItemAction) FilterValue() string {
	return a.Action.Name()
}

func (a ItemAction) Title() string {
	return a.Action.Name()
}

func (a ItemAction) Description() string {
	return a.Action.Description()
}

func FromActionList(l action.List) []ItemAction {
	var actions []ItemAction
	for _, a := range l {
		actions = append(actions, ItemAction{Action: a})
	}
	return actions
}
