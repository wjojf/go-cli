package types

import (
	"github.com/google/uuid"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
	"github.com/wjojf/go-ssh-tui/internal/types/key"
	"golang.org/x/crypto/ssh"
)

// Process is a struct that contains all steps data for the current process
// It is used to pass data between screens(steps in a pipeline)
type Process struct {
	ID               uuid.UUID
	User             string
	Action           action.Action
	Host             host.Host
	PrivateKeyLoader key.PrivateKeyLoader
	Session          *ssh.Session
}

func NewProcess() *Process {
	return &Process{
		ID: uuid.New(),
	}
}
