package keys

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wjojf/go-ssh-tui/internal/screens/docker/monitor"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
	"github.com/wjojf/go-ssh-tui/internal/types/key"
)

type Model struct {
	process *types.Process

	// Keys list
	list list.Model

	// Key files
	keys []key.PrivateKeyLoader

	// Style
	style lipgloss.Style

	// Flow control
	next tea.Model
}

func NewModel(opts ModelOpts) *Model {

	keys := key.GetPrivateKeyLoaders()

	return &Model{
		process: opts.Process,
		keys:    keys,
		list:    GetList(keys),
		next:    nil,
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

	keyLoader := m.keys[m.list.Cursor()]
	m.process.PrivateKeyLoader = keyLoader

	// Set next model
	m.setNext()

	if m.next == nil {
		return m, tea.Quit
	}

	return m.next, m.next.Init()
}

func (m *Model) setNext() {

	var next tea.Model = nil

	switch m.process.Action.ID() {
	case action.DockerMonitorID:
		next = monitor.NewModel(
			monitor.ModelOpts{
				Process: m.process,
			},
		)
	case action.DockerRestartID:
		// TODO: implement models
		next = nil
	case action.SSLRenewID:
		// TODO: implement models
		next = nil
	default:
		next = nil
	}

	m.next = next
}
