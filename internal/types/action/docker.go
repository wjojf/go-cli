package action

const (
	DockerMonitorID = "docker-monitor"
	DockerRestartID = "docker-restart"
)

// DockerMonitor is an action that monitors Docker containers.
type DockerMonitor struct{}

func (d DockerMonitor) ID() string {
	return DockerMonitorID
}

func (d DockerMonitor) Name() string {
	return "Docker Monitor"
}

func (d DockerMonitor) Description() string {
	return "Monitor Docker containers"
}

// DockerRestart is an action that restarts Docker containers.
type DockerRestart struct{}

func (d DockerRestart) ID() string {
	return DockerRestartID
}

func (d DockerRestart) Name() string {
	return "Docker Restart"
}

func (d DockerRestart) Description() string {
	return "Restart Docker containers"
}
