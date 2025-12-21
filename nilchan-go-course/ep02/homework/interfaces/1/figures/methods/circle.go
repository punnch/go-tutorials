package methods

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}
