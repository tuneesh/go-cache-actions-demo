package main

import (
	"log"
	"net/http"
	"os"

	helloapi "example.com/hello-api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("hello API listening on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, helloapi.NewHandler()))
}
