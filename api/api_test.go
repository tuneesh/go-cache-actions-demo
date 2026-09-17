package api

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Version() != "v1" {
		t.Fatalf("Version() = %q, want v1", Version())
	}
}
