package monitor

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"github.com/wjojf/go-ssh-tui/internal/service/docker"
	"github.com/wjojf/go-ssh-tui/internal/types"
)

type Model struct {
	process *types.Process

	loading bool
	err     error

	progress progress.Model

	service *docker.MonitorService
	ingress chan monitorMsg

	table table.Model
	data  []MonitorStatRow
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

		ingress: make(chan monitorMsg),
	}

	m.table = m.refreshTable()

	go m.listen()

	return m
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.startLoading(),
		loaderTick(),
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.shutDown()
			return m, tea.Quit
		case tea.KeyEsc:
			m.shutDown()
			return m, tea.Quit
		}

	// Loading Progress Bar

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, nil
		}

		cmd = m.progress.IncrPercent(0.25)
		return m, tea.Batch(loaderTick(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case errMsg:
		m.loading, m.err = false, msg.Err
		return m, nil

	case loadedMsg:
		m.service, m.loading = msg.Service, false
		m.requestData()
		return m, monitorTick()

	// Monitor Service

	// Request New Data
	case monitorTickMsg:
		m.requestData()
		return m, monitorTick()

	// Monitor Data
	case monitorMsg:
		m.table = m.refreshTable()
		return m, nil
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.loading {
		return m.progress.View()
	}

	if m.err != nil {
		return m.err.Error()
	}

	return baseStyle.Render(m.table.View())
}

func (m *Model) requestData() {
	go func() {
		stats, err := m.service.GetMonitorData()
		m.ingress <- monitorMsg{Stats: stats, Error: err}
	}()
}

func (m *Model) listen() {
	for {
		select {
		case msg := <-m.ingress:
			if msg.Error != nil {
				m.err = msg.Error
				return
			}

			m.data = make([]MonitorStatRow, 0, len(msg.Stats))
			for _, stat := range msg.Stats {
				m.data = append(m.data, MonitorStatRow{Stat: stat})
			}

			// Send a message to update view
			m.Update(msg)
		}
	}
}

func (m *Model) shutDown() {
	m.service.Close()
	close(m.ingress)
}
