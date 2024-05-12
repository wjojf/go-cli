package keys

import (
	"github.com/wjojf/go-ssh-tui/internal/types/key"
)

type KeyPathItem struct {
	Loader key.PrivateKeyLoader
}

func (i KeyPathItem) FilterValue() string {
	return i.Loader.Filepath
}

func (i KeyPathItem) Title() string {
	return i.Loader.Filepath
}

func (i KeyPathItem) Description() string {
	return ""
}
