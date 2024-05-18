package monitor

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/wjojf/go-ssh-tui/internal/service/docker"
)

type MonitorStatRow struct {
	Stat docker.ContainerInfo
}

func (s MonitorStatRow) Row() table.Row {
	return table.Row{
		s.Stat.Names,
		s.Stat.Status,
		s.Stat.CPUUsage,
		s.Stat.MemoryUsage,
		s.Stat.NetworkIO,
	}
}
