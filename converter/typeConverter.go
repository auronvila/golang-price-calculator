package typeConverter

import (
	"bufio"
	"os"
	"strconv"
)

func FileReaderAndParser(file *os.File) ([]float64, error) {
	scanner := bufio.NewScanner(file)
	var lines []float64
	for scanner.Scan() {
		parsedVal, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, parsedVal)
	}

	return lines, nil
}
