package api

import (
	"encoding/json"
	"net/http"

	"github.com/eimiss/functionLib/function"
)

type Handler struct {
	Fn function.Function
}

func NewHandler(fn function.Function) *Handler {
	return &Handler{Fn: fn}
}

func (h *Handler) ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	if input == "" {
		input = "World"
	}

	result := h.Fn.Execute(input)

	json.NewEncoder(w).Encode(map[string]string{
		"result": result,
	})
}
