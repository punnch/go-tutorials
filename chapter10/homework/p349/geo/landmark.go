package geo

import "errors"

// Struct
type Landmark struct {
	name string
	Coordinates
}

// Get name method
func (l *Landmark) Name() string {
	return l.name
}

// Set name method
func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}
