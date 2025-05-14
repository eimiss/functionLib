package main

import (
	"log"
	"net/http"

	"github.com/eimiss/functionLib/api"
	"github.com/eimiss/functionLib/function"
)

func main() {
	fn := &function.ImageToASCIIFunction{}
	handler := api.NewHandler(fn)

	http.HandleFunc("/ascii", handler.ExecuteHandler)

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
