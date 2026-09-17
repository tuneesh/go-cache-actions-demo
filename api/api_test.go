package api

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Version() != "expected-to-fail" {
		t.Fatalf("intentional cache-safety check: Version() = %q, want expected-to-fail", Version())
	}
}
