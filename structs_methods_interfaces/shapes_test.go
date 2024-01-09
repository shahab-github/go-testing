package shapes

import "testing"

// func check(t testing.TB, got, want) {
// 	t.Helper()
// 	if got!= want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
