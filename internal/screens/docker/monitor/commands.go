package monitor

import (
	"context"
	"errors"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/service/docker"
)

func (m *Model) startLoading() tea.Cmd {
	return func() tea.Msg {
		errorsCh := make(chan error, 1)
		outputCh := make(chan *docker.MonitorService, 1)

		ctx, f := context.WithTimeout(context.Background(), time.Second*2)
		defer f()

		go func() {
			defer close(errorsCh)
			defer close(outputCh)

			fmt.Fprintln(m.logger, "Goroutine started, simulating delay...")
			time.Sleep(1 * time.Second)

			service, err := docker.NewMonitorService(
				docker.MonitorOpts{SSHOpts: m.GetSshOpts()},
			)
			if err != nil {
				fmt.Fprintln(m.logger, "Error creating MonitorService:", err)
				errorsCh <- err
				return
			}

			fmt.Fprintln(m.logger, "MonitorService created successfully")
			outputCh <- service
		}()

		for {
			select {
			case err := <-errorsCh:
				fmt.Fprintln(m.logger, "Received error message:", err)
				return errMsg{Err: err}
			case s := <-outputCh:
				fmt.Fprintln(m.logger, "Received loaded message:", s)
				return loadedMsg{Service: s}
			case <-ctx.Done():
				fmt.Fprintln(m.logger, "Context done, timeout")
				return errMsg{Err: errors.New("timeout")}
			}
		}
	}
}
func progressTick() tea.Cmd {
	return func() tea.Msg {
		return tickMsg(time.Now())
	}
}

func immediateTick() tea.Cmd {
	return tea.Batch(
		progressTick(),
		tea.Tick(250*time.Millisecond, func(t time.Time) tea.Msg {
			return tickMsg(t)
		}),
	)
}
