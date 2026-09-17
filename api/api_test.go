package api

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Version() != "version-X" {
		t.Fatalf("Version() = %q, want version-3", Version())
	}
}
