package function

import "fmt"

type EnglishFunction struct{}

func (EnglishFunction) Execute(input string) string {
	return fmt.Sprintf("Hello, %s!", input)
}

type SpanishFunction struct{}

func (SpanishFunction) Execute(input string) string {
	return fmt.Sprintf("¡Hola, %s!", input)
}
