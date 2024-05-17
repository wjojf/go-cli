package monitor

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"github.com/wjojf/go-ssh-tui/internal/service/docker"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"io"
)

type Model struct {
	process *types.Process

	loading bool
	err     error

	progress progress.Model

	service *docker.MonitorService

	logger io.Writer
}

func NewModel(opts ModelOpts) *Model {
	m := &Model{
		process: opts.Process,
		service: nil,

		loading: true,
		err:     nil,

		progress: progress.New(
			progress.WithDefaultGradient(),
			progress.WithColorProfile(termenv.ANSI256),
		),
	}

	return m
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.startLoading(),
		progressTick(),
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m, tea.Quit
		}

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, nil
		}

		cmd = m.progress.IncrPercent(0.25)
		return m, tea.Batch(progressTick(), cmd)

	case errMsg:
		m.loading, m.err = false, msg.Err
		return m, nil

	case loadedMsg:
		m.service, m.loading = msg.Service, false
		return m, nil

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m *Model) View() string {
	if m.loading {
		return m.progress.View()
	}

	if m.err != nil {
		return m.err.Error()
	}

	return "Connected!"
}
