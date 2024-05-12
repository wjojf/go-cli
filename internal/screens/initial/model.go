package initial

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/screens/actions"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/loader"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/logo"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/prompt"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
)

type Model struct {
	// Input
	input textinput.Model

	// Spinner
	spinner spinner.Model

	// Logo
	logo *logo.Logo

	// Flow control
	next tea.Model
}

func NewModel() Model {

	return Model{
		input:   prompt.GetInput(),
		spinner: loader.GetSpinner(),
		logo:    logo.New(logo.GetStyles()),
		next: actions.NewModel(actions.ModelOpts{
			Actions: action.DefaultActions,
			Process: types.NewProcess(),
		}),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		textinput.Blink,
	)
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
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"\n %s \n\n %s %s \n",
		m.logo.Render(),
		m.spinner.View(),
		m.input.View(),
	)
}
