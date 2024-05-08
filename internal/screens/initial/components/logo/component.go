package logo

import (
	"github.com/wjojf/go-ssh-tui/pkg/component"
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

func New() component.Component {
	c := &asciiLogo{}

	c.SetStyle(LogoStyles)

	return c
}

type asciiLogo struct {
	component.Base
}

func (a *asciiLogo) Render() string {
	return a.Style.Render(WelcomeAsciiArt)
}
