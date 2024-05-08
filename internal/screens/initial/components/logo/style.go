package logo

import "github.com/charmbracelet/lipgloss"

var (
	LogoStyles = lipgloss.NewStyle().
		Bold(true).
		// Foreground(lipgloss.Color("F5BA22")).
		BorderStyle(lipgloss.RoundedBorder()).
		// BorderForeground(lipgloss.Color("F5BA22"))
		AlignHorizontal(lipgloss.Center)
)
