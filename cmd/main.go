package main

import (
	"github.com/wjojf/go-bubbletea-sample/internal/pkg/app"
)

func main() {
	app := app.NewApp()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
