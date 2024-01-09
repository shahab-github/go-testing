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

// Refactor to below
// func TestArea(t *testing.T) {

// 	checkArea := func (t testing.TB, shape Shape, want) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got!= want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		checkArea(t, rectangle, 100.0)
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.0)
// 	})
// }

// Again refactor to below
func TestArea(t *testing.T) {

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.0},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
