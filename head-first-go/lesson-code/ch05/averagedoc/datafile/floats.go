// Package datafile reads data from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads data from file and returns []float64
func GetFloats(fileName string) ([]float64, error) {
	numbers := []float64{}

	file, err := os.Open(fileName)
	if err != nil {
		// return nil because slice is nil
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	// While scanner.Scan() true - scan file
	// if file's end or error encountered - false
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// Close file to save RAM
	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
