package ssh

import (
	"io"

	"golang.org/x/crypto/ssh"
)

func NewSession(opts SessionOpts) (*ssh.Session, error) {
	session, err := opts.Client.NewSession()
	if err != nil {
		return nil, err
	}

	session.Stdout = opts.Output
	session.Stderr = opts.Output

	return session, nil
}

type SessionOpts struct {
	Client *ssh.Client
	Output io.Writer
}
