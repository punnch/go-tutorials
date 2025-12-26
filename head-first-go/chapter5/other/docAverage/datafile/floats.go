// Пакет datafile позволяет читать данные с файлов
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из файла и возвращает float64 массив
func GetFloats(fileName string) ([]float64, error) {
	numbers := []float64{}

	file, err := os.Open(fileName)
	if err != nil {
		// возвращаем nil, т.к. пустой слайс ([]float64) == nil
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	// Пока scanner.Scan() true - сканируем, когда ошибка / конец файла - false
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// Закрываем файл, чтобы не засорять память
	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
