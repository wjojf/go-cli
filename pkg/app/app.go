package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wjojf/go-bubbletea-sample/pkg/model"
)

type App struct {
	tea *tea.Program
}

func NewApp() *App {
	return &App{
		tea: tea.NewProgram(
			model.NewShoppingList(model.DefaultItems...),
		),
	}
}

func (a *App) Start() {
	if _, err := a.tea.Run(); err != nil {
		panic(err)
	}
}
