package app

import (
	"fmt"

	"github.com/eimiss/functionLib/function"
)

// The App struct has a single field Fn of type function.Function, which is an interface that defines a function that can be executed.
type App struct {
	Fn function.Function
}

// This is a constructor function in Go that creates a new instance of the App struct.
// It takes a function.Function object as an argument and returns a pointer to a new App instance with the Fn field initialized to the provided function.
func NewApp(fn function.Function) *App {
	return &App{Fn: fn}
}

// Run executes the function.Function object stored in the App instance with the provided parameters, and
// prints the result to the console. If an error occurs during execution, it is printed to the console as
// well.
func (a *App) Run(inputPath string, euclideanDistance int, widthImage int, isColored bool) {
	result, err := a.Fn.Execute(inputPath, euclideanDistance, widthImage, isColored)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
