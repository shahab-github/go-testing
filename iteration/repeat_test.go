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

// Benchmarking
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
