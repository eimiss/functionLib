package main

import (
	"log"
	"net/http"

	"github.com/eimiss/functionLib/api"
	"github.com/eimiss/functionLib/app"
	"github.com/eimiss/functionLib/function"
)

func main() {
	fn := function.EnglishFunction{}
	myApp := app.NewApp(fn)
	myApp.Run("Alice")

	handler := api.NewHandler(fn)
	http.HandleFunc("/execute", handler.ExecuteHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
