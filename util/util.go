package util

import (
	"bufio"
	"io"
	"os"
)

func Read_input(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))

	}
	return lines, nil
}
