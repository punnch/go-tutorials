// package calculator is to calculate numbers

// 1. Available operations:
// +, -, *, /

// 2. Return error if:
// - division by 0
// - arguments are > 1000 or < -1000

// 3. If error encountered, print:
// - operation
// - error

package calculator

import "errors"

type Numbers struct {
	A int
	B int
}

// ArgumentError() checks an available argument range
func argumentError(n Numbers) error {
	if n.A > 1000 || n.A < -1000 || n.B > 1000 || n.B < -1000 {
		return errors.New("invalid arguments")
	}

	return nil
}

// DivisionError() checks in division if the second argument is 0
func divisionError(n Numbers) error {
	if n.B == 0 {
		return errors.New("you can't divide by 0")
	}

	return nil
}

// Sum()
func Sum(n Numbers) (int, error) {
	if err := argumentError(n); err != nil {
		return 0, err
	}

	return n.A + n.B, nil
}

// Subtract()
func Subtract(n Numbers) (int, error) {
	if err := argumentError(n); err != nil {
		return 0, err
	}

	return n.A - n.B, nil
}

// Multiplicate()
func Multiplicate(n Numbers) (int, error) {
	if err := argumentError(n); err != nil {
		return 0, err
	}

	return n.A * n.B, nil
}

// Divide()
func Divide(n Numbers) (int, error) {
	if err := argumentError(n); err != nil {
		return 0, err
	}

	if err := divisionError(n); err != nil {
		return 0, err
	}

	return n.A / n.B, nil
}
