package methods

type Triangle struct {
	A, B, C, H float64
}

func (t Triangle) Area() float64 {
	return (t.H * t.B) / 2
}

func (t Triangle) Perim() float64 {
	return t.H + t.B + t.C
}
