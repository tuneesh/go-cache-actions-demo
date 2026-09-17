package helloapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// NewHandler returns the HTTP routes exposed by the hello API.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /greet", greet)
	mux.HandleFunc("GET /request-id", requestID)
	return mux
}

func home(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"service": "hello-api"})
}

func health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	message := "hello, " + name
	if r.URL.Query().Get("reverse") == "true" {
		message = reverse(message)
	}
	if r.URL.Query().Get("shout") == "true" {
		message = strings.ToUpper(message)
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": message})
}

func requestID(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"request_id": uuid.NewString()})
}

func reverse(value string) string {
	runes := []rune(value)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
