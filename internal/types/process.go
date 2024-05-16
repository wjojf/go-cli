package types

import (
	"github.com/wjojf/go-ssh-tui/internal/types/action"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
	"github.com/wjojf/go-ssh-tui/internal/types/key"
)

// Process is a struct that contains all steps data for the current process
// It is used to pass data between screens(steps in a pipeline)
type Process struct {
	User             string
	Action           action.Action
	Host             *host.Host
	PrivateKeyLoader *key.PrivateKeyLoader
}

func NewProcess() *Process {
	return &Process{
		User: "ubuntu",
	}
}
