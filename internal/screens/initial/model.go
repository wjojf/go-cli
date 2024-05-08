package initial

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/pkg/component"
)

type Model struct {
	user    string           // Welcome Message
	actions types.ActionList // Possible actions
	tree    component.Tree

	// Flow control
	next *tea.Model
}

func NewModel(opts ModelOpts) Model {
	return Model{
		user:    opts.User,
		actions: opts.Actions,
		tree:    components.GetInititalTree(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m Model) View() string {
	return m.tree.Render()
}
