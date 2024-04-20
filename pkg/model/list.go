package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ShoppingList struct {
	items    []*ShoppingItem
	cursor   int
	selected map[int]struct{}
}

func NewShoppingList(items ...*ShoppingItem) *ShoppingList {
	model := &ShoppingList{
		items:    make([]*ShoppingItem, 0),
		selected: make(map[int]struct{}),
	}

	for _, item := range items {
		model.AddItem(item)
	}

	return model
}

func (sl *ShoppingList) AddItem(item *ShoppingItem) {
	sl.items = append(sl.items, item)
}

func (sl *ShoppingList) Init() tea.Cmd {
	return nil
}

func (sl *ShoppingList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return sl, tea.Quit

		case "up", "k":
			sl.Up()

		case "down", "j":
			sl.Down()

		case "enter", " ":
			sl.Toggle(sl.cursor)

		}
	}

	return sl, nil
}

func (sl *ShoppingList) View() string {
	s := "Shopping List\n\n"

	for i, item := range sl.items {

		cursor := " "
		if sl.cursor == i {
			cursor = ">"
		}

		selected := " "
		if _, ok := sl.selected[i]; ok {
			selected = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, item.Name())

	}

	s += "\nPress q to quit.\n"

	return s
}

func (sl *ShoppingList) Toggle(i int) {
	if _, ok := sl.selected[i]; ok {
		delete(sl.selected, i)
		return
	}
	sl.selected[i] = struct{}{}
}

func (sl *ShoppingList) Up() {
	if sl.cursor > 0 {
		sl.cursor--
	}
}

func (sl *ShoppingList) Down() {
	if sl.cursor < len(sl.items)-1 {
		sl.cursor++
	}
}
