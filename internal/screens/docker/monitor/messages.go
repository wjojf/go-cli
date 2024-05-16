package monitor

import "github.com/wjojf/go-ssh-tui/internal/service/docker"

type errMsg struct {
	Err error
}

type loadedMsg struct {
	Service *docker.MonitorService
}

type tickMsg struct{}
