package main

import (
	"fmt"
	"log"

	"github.com/eimiss/functionLib/function"
)

// This example shows how to use the ImageToASCIIFunction struct to convert various images to ASCII art.
func main() {
	fn := function.ImageToASCIIFunction{}
	result, err := fn.Execute("cat.png", 200000, 50, true) // or path to any image
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = fn.Execute("tulip.jpg", 100000, 50, true) // or path to any image
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = fn.Execute("tulip.jpg", 100000, 50, false) // or path to any image
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = fn.Execute("bear.png", 50000, 50, true) // or path to any image
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
