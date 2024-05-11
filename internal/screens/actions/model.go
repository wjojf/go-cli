package actions

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wjojf/go-ssh-tui/internal/screens/hosts"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
)

type Model struct {
	actions []ItemAction
	process *types.Process

	list  list.Model
	style lipgloss.Style

	next tea.Model
}

func NewModel(opts ModelOpts) *Model {

	actions := FromActionList(opts.Actions)

	l := GetList(actions...)

	return &Model{
		process: opts.Process,
		actions: actions,
		style:   GetListStyles(80),
		list:    l,
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
		case tea.KeyEnter:
			return m.handleEnter()
		}
	case tea.WindowSizeMsg:
		h, v := m.style.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.style = GetListStyles(msg.Width)
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	return m.style.Render(m.list.View())
}

func (m *Model) handleEnter() (tea.Model, tea.Cmd) {

	// Get the selected item
	action := m.actions[m.list.Cursor()].Action

	// Set the action in the process
	m.process.Action = action

	return hosts.NewModel(hosts.ModelOpts{
		Hosts:   host.AllHosts,
		Process: m.process,
	}), nil
}
