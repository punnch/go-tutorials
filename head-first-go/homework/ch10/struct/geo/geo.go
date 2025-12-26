// package geo need to work with coordinates
package geo

import "errors"

// Struct
type Coordinates struct {
	latitude  float64
	longitude float64
}

// Get Latitude
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

// Get Longitude
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

// Set Latitude
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

// Set Longitude
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}
