package keys

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/wjojf/go-ssh-tui/internal/types/key"
)

func GetList(keys []key.PrivateKeyLoader) list.Model {

	var items []list.Item
	for _, k := range keys {
		items = append(items, KeyPathItem{k})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a key file"

	return l
}
