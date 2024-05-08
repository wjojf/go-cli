package components

import (
	"github.com/wjojf/go-ssh-tui/internal/screens/initial/components/logo"
	"github.com/wjojf/go-ssh-tui/pkg/component"
)

func GetInititalTree() component.Tree {

	// Header
	logo := logo.New()

	return component.NewTree(logo)
}
