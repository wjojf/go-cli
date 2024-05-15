package docker

import (
	client "github.com/wjojf/go-ssh-tui/pkg/ssh"
	"golang.org/x/crypto/ssh"
)

type MonitorService struct {
	session *ssh.Session
}

func NewMonitorService(opts MonitorOpts) (*MonitorService, error) {

	c, err := client.NewClient(opts.SSHOpts)
	if err != nil {
		return nil, err
	}

	session, err := c.NewSession()
	if err != nil {
		return nil, err
	}

	return &MonitorService{
		session: session,
	}, nil
}
