package prompt

import "github.com/charmbracelet/bubbles/textinput"

func GetInput() textinput.Model {

	ti := textinput.New()
	ti.Placeholder = "Press Enter to continue... (Ctrl+C to quit)"

	ti.Focus()

	ti.CharLimit = 156
	ti.Width = 20

	return ti
}
