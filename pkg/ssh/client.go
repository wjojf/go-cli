package ssh

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

func NewClient(opts *SSHOpts) (*ssh.Client, error) {
	return ssh.Dial(opts.Protocol, opts.Addr, opts.CfgBuilder.Build())
}

func ExecuteCommand(session *ssh.Session, cmd string) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return "", err
	}
	return b.String(), nil
}
