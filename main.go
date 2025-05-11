package main

import (
	"log"
	"net/http"

	"github.com/eimiss/library/api"
	"github.com/eimiss/library/app"
	"github.com/eimiss/library/function"
)

func main() {
	// Command-line run
	myGreeter := greeter.EnglishGreeter{}
	myApp := app.NewApp(myGreeter)
	myApp.Run("Alice")

	// Start HTTP server
	handler := api.NewHandler(myGreeter)
	http.HandleFunc("/greet", handler.GreetHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
