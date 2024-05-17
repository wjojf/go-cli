package hosts

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wjojf/go-ssh-tui/internal/screens/keys"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
)

type Model struct {
	process *types.Process

	hosts []host.Host

	list  list.Model
	style lipgloss.Style

	next tea.Model
}

func NewModel(opts ModelOpts) *Model {

	l := GetList(opts.Hosts...)

	return &Model{
		process: opts.Process,
		hosts:   opts.Hosts,
		style:   GetListStyles(80),
		list:    l,
		next: keys.NewModel(keys.ModelOpts{
			Process: opts.Process,
		}),
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
	if m.process.Action == nil {
		return "No action selected"
	}

	var s string

	s += m.style.Render(m.list.View())

	return s
}

func (m *Model) handleEnter() (tea.Model, tea.Cmd) {

	host := m.hosts[m.list.Cursor()]

	m.process.Host = host

	return m.next, nil
}
