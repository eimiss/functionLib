package app

import (
	"fmt"

	"github.com/eimiss/library/function"
)

type App struct {
	Greeter function.Greeter
}

func NewApp(g function.Greeter) *App {
	return &App{Greeter: g}
}

func (a *App) Run(name string) {
	msg := a.Greeter.Greet(name)
	fmt.Println(msg)
}
