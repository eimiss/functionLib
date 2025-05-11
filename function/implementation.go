package greeter

import "fmt"

type EnglishGreeter struct{}

func (EnglishGreeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type SpanishGreeter struct{}

func (SpanishGreeter) Greet(name string) string {
	return fmt.Sprintf("¡Hola, %s!", name)
}
