package calendar

import "errors"

// Struct
type Date struct {
	year  int
	month int
	day   int
}

// Set Year method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// Set Month method
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// Set Day method
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// Get Year method
func (d *Date) Year() int {
	return d.year
}

// Get Month method
func (d *Date) Month() int {
	return d.month
}

// Get Day method
func (d *Date) Day() int {
	return d.day
}
