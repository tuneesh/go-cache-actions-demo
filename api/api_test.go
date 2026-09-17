package api

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Version() != "version-3" {
		t.Fatalf("Version() = %q, want version-3", Version())
	}
}
