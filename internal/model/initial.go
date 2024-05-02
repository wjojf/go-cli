package model

import tea "github.com/charmbracelet/bubbletea"

type InitialModel struct{}

func NewInitialModel() InitialModel {
	return InitialModel{}
}

func (m InitialModel) Init() tea.Cmd {
	return nil
}

func (m InitialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m InitialModel) View() string {
	return "TODO"
}
