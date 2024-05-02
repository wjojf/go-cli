package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-bubbletea-sample/internal/model"
	"github.com/wjojf/go-bubbletea-sample/internal/pkg/config"
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
			model.NewInitialModel(),
		),
	}
}

func (a *App) Start() error {
	if _, err := a.tea.Run(); err != nil {
		return err
	}

	return nil
}
