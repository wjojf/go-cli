package ssh

import "golang.org/x/crypto/ssh"

var (
	DefaultSSHOpts = &SSHOpts{
		CfgBuilder: NewConfigBuilder(),
		Protocol:   "tcp",
		Addr:       "localhost",
	}
)

type SSHOpts struct {
	CfgBuilder *ConfigBuilder
	Protocol   string
	Addr       string
}

type ConfigBuilder struct {
	cfg *ssh.ClientConfig
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		cfg: &ssh.ClientConfig{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}
}

func (b *ConfigBuilder) Build() *ssh.ClientConfig {
	return b.cfg
}

func (b *ConfigBuilder) WithUser(user string) *ConfigBuilder {
	b.cfg.User = user
	return b
}

func (b *ConfigBuilder) WithPrivateKey(privateKey string) *ConfigBuilder {

	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return b
	}

	b.cfg.Auth = append(b.cfg.Auth, ssh.PublicKeys(key))

	return b
}

func (b *ConfigBuilder) WithHostKeyCallback(callback ssh.HostKeyCallback) *ConfigBuilder {
	b.cfg.HostKeyCallback = callback
	return b
}
