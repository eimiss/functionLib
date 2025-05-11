package api

import (
	"encoding/json"
	"net/http"

	"github.com/eimiss/library/function"
)

type Handler struct {
	Greeter function.Greeter
}

func NewHandler(g function.Greeter) *Handler {
	return &Handler{Greeter: g}
}

func (h *Handler) GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	message := h.Greeter.Greet(name)

	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
