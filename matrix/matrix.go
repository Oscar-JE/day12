package matrix

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

type Matrix struct {
	arr [][]int
}

func Init() Matrix {
	return Matrix{[][]int{}}
}

func Zeros(nrRows int, nrCols int) Matrix {
	return FilledWith(nrRows, nrCols, 0)
}

func FilledWith(nrRows int, nrCols int, value int) Matrix {
	array := [][]int{}
	i := 0
	for i < nrRows {
		j := 0
		row := []int{}
		for j < nrCols {
			row = append(row, value)
			j++
		}
		array = append(array, row)
		i++
	}
	return Matrix{arr: array}
}

func (m Matrix) Get(row int, col int) int {
	return m.arr[row][col]
}

func (m *Matrix) Set(row int, col int, value int) {
	m.arr[row][col] = value
}

func (m Matrix) String() string {
	rep := ""
	for _, row := range m.arr {
		for numIndex := range row {
			numRep := fmt.Sprint(row[numIndex])
			for len(numRep) < 4 {
				numRep += " "
			}
			rep += numRep
		}
		rep = rep + "\n"
	}
	return rep
}

func (m *Matrix) Append(row []int) {
	m.arr = append(m.arr, row)
}

func (m Matrix) GetNrCols() int {
	return len(m.arr[0])
}

func (m Matrix) GetNrRows() int {
	return len(m.arr)
}

func (m Matrix) InMatrix(p Point) bool {
	inRows := 0 <= p.X && p.X < m.GetNrRows()
	inCols := 0 <= p.Y && p.Y < m.GetNrCols()
	return inRows && inCols
}

func (m Matrix) FindFirst(element int) Point {
	nrRows := len(m.arr)
	nrCols := len(m.arr[0])
	i := 0
	for i < nrRows {
		j := 0
		for j < nrCols {
			if element == m.Get(i, j) {
				return Point{X: i, Y: j}
			}
			j++
		}
		i++
	}
	return Point{0, 0}
}
