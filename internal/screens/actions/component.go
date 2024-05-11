package actions

import "github.com/charmbracelet/bubbles/list"

func GetList(actions ...ItemAction) list.Model {

	var items []list.Item
	for _, a := range actions {
		items = append(items, a)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Choose an action (press 'Enter' to select)"

	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)

	return l
}
