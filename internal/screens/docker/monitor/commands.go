package monitor

import (
	"context"
	"errors"
	"time"

	"github.com/wjojf/go-ssh-tui/internal/service/docker"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) startLoading() tea.Cmd {
	return func() tea.Msg {
		errorsCh := make(chan error, 1)
		outputCh := make(chan *docker.MonitorService, 1)

		ctx, f := context.WithTimeout(context.Background(), time.Second*5)
		defer f()

		go func() {
			defer close(errorsCh)
			defer close(outputCh)

			time.Sleep(1 * time.Second)

			service, err := docker.NewMonitorService(
				docker.MonitorOpts{SSHOpts: m.GetSshOpts()},
			)
			if err != nil {
				errorsCh <- err
				return
			}

			outputCh <- service
		}()

		for {
			select {
			case err := <-errorsCh:
				return errMsg{Err: err}
			case s := <-outputCh:
				return loadedMsg{Service: s}
			case <-ctx.Done():
				return errMsg{Err: errors.New("timeout")}
			}
		}
	}
}
func loaderTick() tea.Cmd {
	return tea.Tick(250*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func monitorTick() tea.Cmd {
	return tea.Tick(10*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
