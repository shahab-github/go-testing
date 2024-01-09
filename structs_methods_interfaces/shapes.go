package shapes

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.width
}
