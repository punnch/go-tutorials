// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat gets an input value and returns float64 type or error encountered
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// get off \n to convert str to float64
	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
