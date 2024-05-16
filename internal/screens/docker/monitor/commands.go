package monitor

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func progressTick() tea.Cmd {
	return tea.Tick(250*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
