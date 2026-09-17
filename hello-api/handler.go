package helloapi

import (
	"encoding/json"
	"net/http"
)

// NewHandler returns the HTTP routes exposed by the hello API.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /greet", greet)
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

	writeJSON(w, http.StatusOK, map[string]string{"message": "hello, " + name})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
