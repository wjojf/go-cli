package initial

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/logo"
	"github.com/wjojf/go-ssh-tui/internal/types"
)

type Model struct {
	actions types.ActionList // Possible actions

	// Flow control
	next *tea.Model
}

func NewModel(opts ModelOpts) Model {
	return Model{
		actions: opts.Actions,
	}
}

func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var s string

	l := logo.New(logo.GetStyles())
	s += l.Render()

	// TODO: add input

	return s
}
