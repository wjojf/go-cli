package actions

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Model struct {
	actions []FilterableAction

	list  list.Model
	style lipgloss.Style
}

func NewModel(opts ModelOpts) *Model {

	actions := FromActionList(opts.Actions)

	return &Model{
		actions: actions,
		style:   ListStyle,
		list:    GetList(actions...),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		log.Printf("Window size changed: %d x %d", msg.Width, msg.Height)
		h, v := m.style.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	return m.style.Render(m.list.View())
}
