package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-ssh-tui/internal/pkg/config"
	"github.com/wjojf/go-ssh-tui/internal/screens/initial"
)

type App struct {
	tea *tea.Program
}

func NewApp() *App {

	err := config.SetupFromEnv()
	if err != nil {
		panic(err)
	}

	return &App{
		tea: tea.NewProgram(
			initial.NewModel(),
		),
	}
}

func (a *App) Start() error {
	if _, err := a.tea.Run(); err != nil {
		return err
	}

	return nil
}
