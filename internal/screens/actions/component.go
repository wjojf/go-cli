package actions

import "github.com/charmbracelet/bubbles/list"

func GetList(actions ...FilterableAction) list.Model {

	var items []list.Item
	for _, a := range actions {
		items = append(items, a)
	}

	l := list.New(items, list.NewDefaultDelegate(), 25, 25)
	l.Title = "Choose an action"

	return l
}
