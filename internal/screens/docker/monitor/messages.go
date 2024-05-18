package monitor

import (
	"time"

	"github.com/wjojf/go-ssh-tui/internal/service/docker"
)

type errMsg struct {
	Err error
}

type loadedMsg struct {
	Service *docker.MonitorService
}

type monitorMsg struct {
	Stats []docker.ContainerInfo
	Error error
}

type tickMsg time.Time

type monitorTickMsg time.Time
