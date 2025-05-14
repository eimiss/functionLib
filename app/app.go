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

func (a *App) Run(inputPath string, euclideanDistance int, widthImage int, isColored bool) {
	result, err := a.Fn.Execute(inputPath, euclideanDistance, widthImage, isColored)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
