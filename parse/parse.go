package parse

import (
	"bufio"
	"day12/matrix"
	"os"
)

func parse(filePath string) matrix.Matrix {
	file, err := os.Open(filePath)
	if err == nil {
		reader := bufio.NewScanner(file)
		return parseMatrix(reader)
	}
	panic(err)
}

func parseMatrix(scanner *bufio.Scanner) matrix.Matrix {
	line := ""
	if scanner.Scan() {
		line = scanner.Text()
	}
	m := matrix.Init(len(line))
	row := convert(line)
	m.Append(row)
	return m
}

func convert(line string) []int {
	row := []int{}
	for _, char := range line {
		row = append(row, convertToHeight(char))
	}
	return row
}

func convertToHeight(char rune) int {
	return int(char) - int('a')
}
