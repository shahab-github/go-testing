package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected '%s', but '%s'", expected, got)
	}
}
