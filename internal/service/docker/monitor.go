package docker

import (
	"encoding/json"
	"fmt"
	"strings"

	client "github.com/wjojf/go-ssh-tui/pkg/ssh"
	"golang.org/x/crypto/ssh"
)

type MonitorService struct {
	client *ssh.Client

	stats map[string]DockerStatsResult
	ps    map[string]DockerPsResult
}

func NewMonitorService(opts MonitorOpts) (*MonitorService, error) {

	c, err := client.NewClient(opts.SSHOpts)
	if err != nil {
		return nil, err
	}

	return &MonitorService{
		client: c,
		stats:  make(map[string]DockerStatsResult),
		ps:     make(map[string]DockerPsResult),
	}, nil
}

func (m *MonitorService) Close() error {
	return m.client.Close()
}

func (m *MonitorService) GetMonitorData() ([]ContainerInfo, error) {

	if err := m.getDockerStats(); err != nil {
		return nil, err
	}

	if err := m.getDockerPs(); err != nil {
		return nil, err
	}

	var containers = make([]ContainerInfo, 0, len(m.stats))
	for id, stat := range m.stats {
		if ps, ok := m.ps[id]; ok {
			containers = append(containers, GetDockerContainerInfo(ps, stat))
		}
	}

	return containers, nil
}

func (m *MonitorService) getDockerStats() error {

	session, err := m.client.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	out, err := client.ExecuteCommand(session, "docker stats --no-stream --format '{{json .}}'")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		var result DockerStatsResult
		if err := json.Unmarshal([]byte(line), &result); err != nil {
			return fmt.Errorf("failed to unmarshal 'docker stats' output: %v", err)
		}

		m.stats[result.ID] = result
	}

	return nil
}

func (m *MonitorService) getDockerPs() error {

	session, err := m.client.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	out, err := client.ExecuteCommand(session, "docker ps --format '{{json .}}'")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		var result DockerPsResult
		if err := json.Unmarshal([]byte(line), &result); err != nil {
			return fmt.Errorf("failed to unmarshal 'docker ps' output: %v", err)
		}

		m.ps[result.ID] = result
	}

	return nil
}
