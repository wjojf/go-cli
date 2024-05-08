package component

import "github.com/charmbracelet/lipgloss"

type Base struct {
	Style lipgloss.Style
}

func (b Base) SetStyle(style lipgloss.Style) {
	b.Style = style
}
