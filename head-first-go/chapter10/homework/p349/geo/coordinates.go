package geo

import "errors"

// Struct
type Coordinates struct {
	latitude  float64
	longitude float64
}

// Get latitude method
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

// Get longitude method
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

// Set latitude method
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

// Set longitude method
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}
