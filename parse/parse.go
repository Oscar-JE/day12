package parse

import (
	"bufio"
	"day12/matrix"
	"os"
)

func Parse(filePath string) (matrix.Matrix, matrix.Point, matrix.Point) {
	file, err := os.Open(filePath)
	var m matrix.Matrix
	if err == nil {
		reader := bufio.NewScanner(file)
		m = parseMatrix(reader)
		start := m.FindFirst(ConvertToHeight('S'))
		errStart := m.Set(start.X, start.Y, ConvertToHeight('a'))
		end := m.FindFirst(ConvertToHeight('E'))
		errEnd := m.Set(end.X, end.Y, ConvertToHeight('z'))
		if errStart != nil {
			panic(errStart)
		}
		if errEnd != nil {
			panic(errEnd)
		}
		return m, start, end
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
	for scanner.Scan() {
		line = scanner.Text()
		row := convert(line)
		m.Append(row)
	}
	return m
}

func convert(line string) []int {
	row := []int{}
	for _, char := range line {
		row = append(row, ConvertToHeight(char))
	}
	return row
}

func ConvertToHeight(char rune) int {
	return int(char) - int('a')
}
