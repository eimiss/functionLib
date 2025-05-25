package api

import (
	"net/http"
	"strconv"

	"github.com/eimiss/functionLib/function"
)

// The Handler struct represents an HTTP handler that wraps a function.Function object.
type Handler struct {
	Fn function.Function
}

// NewHandler(fn function.Function) *Handler: Creates a new Handler instance with the given function.Function object.
func NewHandler(fn function.Function) *Handler {
	return &Handler{Fn: fn}
}

// (h *Handler) ExecuteHandler(w http.ResponseWriter, r *http.Request):
// Handles an HTTP request by executing the wrapped function.Function object with query parameters from the request.
func (h *Handler) ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	input := r.URL.Query().Get("input")
	distanceStr := r.URL.Query().Get("distance")
	widthStr := r.URL.Query().Get("width")
	coloredStr := r.URL.Query().Get("colored")

	if input == "" {
		http.Error(w, "'input' parameter is required", http.StatusBadRequest)
		return
	}

	// Convert string params to proper types
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		http.Error(w, "'distance' must be an integer", http.StatusBadRequest)
		return
	}
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		http.Error(w, "'width' must be an integer", http.StatusBadRequest)
		return
	}
	colored, err := strconv.ParseBool(coloredStr)
	if err != nil {
		http.Error(w, "'colored' must be true or false", http.StatusBadRequest)
		return
	}

	// Run the ASCII function
	result, err := h.Fn.Execute(input, distance, width, colored)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send raw colored text output
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
