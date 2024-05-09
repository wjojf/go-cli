package logo

import "github.com/charmbracelet/lipgloss"

func GetStyles() lipgloss.Style {

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		Bold(true).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#FF00FF"))
}
