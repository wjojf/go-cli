package actions

import (
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
)

type ModelOpts struct {
	Actions action.List
	Process *types.Process
}
