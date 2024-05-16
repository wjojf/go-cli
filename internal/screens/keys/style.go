package keys

import "github.com/charmbracelet/lipgloss"

func GetListStyles(width int) lipgloss.Style {
	style := lipgloss.NewStyle().
		Width(width).
		AlignHorizontal(lipgloss.Center)

	return style
}
