package hosts

import (
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
)

type ModelOpts struct {
	Hosts   []host.Host
	Process *types.Process
}
