package billing

import (
	"testing"
	"time"
)

func TestCurrency(t *testing.T) {
	time.Sleep(2 * time.Second)
	if Currency() != "USD" {
		t.Fatalf("Currency() = %q, want USD", Currency())
	}
}
