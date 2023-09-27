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
		m.Set(start.X, start.Y, ConvertToHeight('a'))
		end := m.FindFirst(ConvertToHeight('E'))
		m.Set(end.X, end.Y, ConvertToHeight('z'))
		return m, start, end
	}

	panic(err)
}

func parseMatrix(scanner *bufio.Scanner) matrix.Matrix {
	line := ""
	if scanner.Scan() {
		line = scanner.Text()
	}
	m := matrix.Init()
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
	runes := []rune(line)
	row := []int{}
	j := 0
	for j < len(line) {
		row = append(row, ConvertToHeight(runes[j]))
		j++
	}
	return row
}

func ConvertToHeight(char rune) int {
	heights := "abcdefghijklmnopqrstuvwxyz"
	h := 0
	for _, r := range heights {
		if r == char {
			break
		}
		h++
	}
	return h
}
