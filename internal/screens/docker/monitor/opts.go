package monitor

import (
	"github.com/wjojf/go-ssh-tui/internal/types"
	client "github.com/wjojf/go-ssh-tui/pkg/ssh"
)

type ModelOpts struct {
	Process *types.Process
}

func (m *Model) GetSshOpts() *client.SSHOpts {
	opts := &client.SSHOpts{}

	key, err := m.process.PrivateKeyLoader.Load()
	if err != nil {
		return opts
	}

	cfg := client.NewConfigBuilder().
		WithPrivateKey(string(key)).
		WithUser(m.process.User)

	addr, ok := m.process.Host.IP()
	if !ok {
		return opts
	}

	addr += ":22"

	opts.CfgBuilder = cfg
	opts.Protocol = "tcp"
	opts.Addr = addr

	return opts
}
