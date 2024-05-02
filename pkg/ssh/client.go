package ssh

import "golang.org/x/crypto/ssh"

func NewClient(opts *SSHOpts) (*ssh.Client, error) {
	return ssh.Dial(opts.Protocol, opts.Addr, opts.CfgBuilder.Build())
}
