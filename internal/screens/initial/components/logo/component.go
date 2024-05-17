package logo

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	WelcomeAsciiArt = `
.____     ___________  ________ ________   _________  _____.___.
|    |    \_   _____/ /  _____/ \_____  \  \_   ___ \ \__  |   |
|    |     |    __)_ /   \  ___  /   |   \ /    \  \/  /   |   |
|    |___  |        \\    \_\  \/    |    \\     \____ \____   |
|_______ \/_______  / \______  /\_______  / \______  / / ______|
        \/        \/         \/         \/         \/  \/        `
)

func New(style lipgloss.Style) *Logo {
	return &Logo{
		Style: style,
	}
}

type Logo struct {
	Style lipgloss.Style
}

func (a *Logo) Render() string {

	return a.Style.Render(WelcomeAsciiArt)
}
