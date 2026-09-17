package helloapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health", nil))

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if body := response.Body.String(); body != "{\"status\":\"ok\"}\n" {
		t.Fatalf("body = %q, want health JSON", body)
	}
}

func TestHome(t *testing.T) {
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if body := response.Body.String(); body != "{\"service\":\"hello-api\"}\n" {
		t.Fatalf("body = %q, want service JSON", body)
	}
}

func TestGreet(t *testing.T) {
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/greet?name=Tuneesh", nil))

	if body := response.Body.String(); body != "{\"message\":\"hello, Tuneesh\"}\n" {
		t.Fatalf("body = %q, want greeting JSON", body)
	}
}

func TestGreetShout(t *testing.T) {
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/greet?name=Tuneesh&shout=true", nil))

	if body := response.Body.String(); body != "{\"message\":\"HELLO, TUNEESH\"}\n" {
		t.Fatalf("body = %q, want uppercase greeting JSON", body)
	}
}
