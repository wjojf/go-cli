package initial

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/screens/actions"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/logo"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/prompt"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
)

type Model struct {
	// Input
	input textinput.Model

	// Flow control
	next tea.Model
}

func NewModel() Model {
	return Model{
		input: prompt.GetInput(),
		next: actions.NewModel(actions.ModelOpts{
			Actions: action.DefaultActions,
			Process: types.NewProcess(),
		}),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.next == nil {
				return m, tea.Quit
			}
			return m.next, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var s string

	l := logo.New(logo.GetStyles())
	s += l.Render() + "\n"

	s += m.input.View()

	return s
}
