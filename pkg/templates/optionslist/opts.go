package optionslist

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/wjojf/go-ssh-tui/pkg/utils"
)

type Opts[T utils.Stringable] struct {
	Options []T
	Style   lipgloss.Style
}
