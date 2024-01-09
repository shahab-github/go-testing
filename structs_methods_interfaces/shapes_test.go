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
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	got := Area(rectangle)
// 	want := 72.0

// 	if got != want {
// 		t.Errorf("got %.2f, want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.0

		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	})
}
