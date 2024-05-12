package keys

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"github.com/wjojf/go-ssh-tui/internal/types/key"
)

type Model struct {
	process *types.Process

	// Keys list
	list list.Model

	// Key files
	keys []key.PrivateKeyLoader

	// Flow control
	next tea.Model
}

func NewModel() *Model {
	return &Model{}
}
