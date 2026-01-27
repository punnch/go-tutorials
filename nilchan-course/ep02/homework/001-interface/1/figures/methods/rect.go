package methods

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perim() float64 {
	return 2 * (r.Width + r.Height)
}
