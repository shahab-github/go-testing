package shapes

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
