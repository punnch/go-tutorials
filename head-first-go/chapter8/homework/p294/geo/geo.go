package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Landmark struct {
	Name string
	Coordinates
}
