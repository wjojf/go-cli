package monitor

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/service/docker"
	"github.com/wjojf/go-ssh-tui/internal/types"
)

type Model struct {
	process *types.Process

	loading bool
	err     error

	progress progress.Model

	service *docker.MonitorService
}

func NewModel(opts ModelOpts) *Model {
	return &Model{
		process: opts.Process,
		service: nil,

		loading: true,
		err:     nil,

		progress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		progressTick(),
		m.startLoading,
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg.Err
		return m, nil

	case loadedMsg:
		m.service = msg.Service
		m.loading = false
		return m, nil

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		cmd := m.progress.IncrPercent(0.25)

		return m, tea.Batch(progressTick(), cmd)
	}

	return m, cmd
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

func (m *Model) startLoading() tea.Msg {
	errors := make(chan error, 1)
	output := make(chan *docker.MonitorService, 1)

	go func() {
		defer close(errors)
		defer close(output)

		serice, err := docker.NewMonitorService(
			docker.MonitorOpts{SSHOpts: m.GetSshOpts()},
		)
		if err != nil {
			errors <- err
			return
		}

		output <- serice
	}()

	// Artificially slow down the loading process so we can see the progress bar
	time.Sleep(1 * time.Second)

	for {
		select {
		case err := <-errors:
			return errMsg{Err: err}
		case s := <-output:
			return loadedMsg{Service: s}
		}
	}
}
