package app

import (
	"fmt"

	"github.com/eimiss/functionLib/function"
)

type App struct {
	Fn function.Function
}

func NewApp(fn function.Function) *App {
	return &App{Fn: fn}
}

func (a *App) Run(input string) {
	result := a.Fn.Execute(input)
	fmt.Println(result)
}
