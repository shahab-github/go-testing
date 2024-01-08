package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but '%d'", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
