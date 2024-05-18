package docker

import (
	"fmt"
	"testing"

	"github.com/wjojf/go-ssh-tui/internal/types/key"
	client "github.com/wjojf/go-ssh-tui/pkg/ssh"
)

func TestMonitoringService(t *testing.T) {

	keyLoader := key.GetPrivateKeyLoaders()[2]
	key, err := keyLoader.Load()
	if err != nil {
		t.Errorf("Error loading private key: %v", err)
	}

	cfgBuilder := client.NewConfigBuilder().
		WithUser("ubuntu").
		WithPrivateKey(string(key))

	opts := MonitorOpts{
		SSHOpts: &client.SSHOpts{
			CfgBuilder: cfgBuilder,
			Addr:       "185.224.139.29:22",
			Protocol:   "tcp",
		},
	}

	s, err := NewMonitorService(opts)
	if err != nil {
		t.Errorf("Error creating monitor service: %v", err)
	}

	data, err := s.GetMonitorData()
	if err != nil {
		t.Errorf("Error getting monitor data: %v", err)
	}

	if len(data) == 0 {
		t.Errorf("No data returned")
	}

	for _, d := range data {
		fmt.Printf("%+v", d)
	}
}
