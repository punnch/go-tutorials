// Пакет datafile нужен для чтения файлов
package datafile

import (
	"bufio"
	"os"
)

// GetStrings читает каждую строку из .txt файла
func GetStrings(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		vote := scanner.Text()
		lines = append(lines, vote)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
