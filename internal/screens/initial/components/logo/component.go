package logo

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	WelcomeAsciiArt = `
.____     ___________  ________ ________   _________  _____.___.   _________  ___ ___  ___________.____     .____     
|    |    \_   _____/ /  _____/ \_____  \  \_   ___ \ \__  |   |  /   _____/ /   |   \ \_   _____/|    |    |    |    
|    |     |    __)_ /   \  ___  /   |   \ /    \  \/  /   |   |  \_____  \ /    ~    \ |    __)_ |    |    |    |    
|    |___  |        \\    \_\  \/    |    \\     \____ \____   |  /        \\    Y    / |        \|    |___ |    |___ 
|_______ \/_______  / \______  /\_______  / \______  / / ______| /_______  / \___|_  / /_______  /|_______ \|_______ \
        \/        \/         \/         \/         \/  \/                \/        \/          \/         \/        \/
`
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
