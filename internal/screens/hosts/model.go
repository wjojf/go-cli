package hosts

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
)

type Model struct {
	process *types.Process

	hosts []host.Host

	list  list.Model
	style lipgloss.Style
}

func NewModel(opts ModelOpts) *Model {

	l := GetList(opts.Hosts...)

	return &Model{
		process: opts.Process,
		hosts:   opts.Hosts,
		style:   GetListStyles(80),
		list:    l,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) View() string {
	if m.process.Action == nil {
		return "No action selected"
	}

	return m.process.Action.Name()
}
