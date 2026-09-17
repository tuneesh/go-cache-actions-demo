package database

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Name() != "primary" {
		t.Fatalf("Name() = %q, want primary", Name())
	}
}
