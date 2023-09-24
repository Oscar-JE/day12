package matrix

import (
	"errors"
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
	arr    []int
	nrRows int
	nrCols int
}

func Init(nrCols int) Matrix {
	return Matrix{[]int{}, 0, nrCols}
}

func Zeros(nrRows int, nrCols int) Matrix {
	return FilledWith(nrRows, nrCols, 0)
}

func FilledWith(nrRows int, nrCols int, value int) Matrix {
	array := []int{}
	for i := 0; i < nrRows*nrCols; i++ {
		array = append(array, value)
	}
	return Matrix{arr: array, nrRows: nrRows, nrCols: nrCols}
}

func (m *Matrix) Append(row []int) error {
	if len(row) == m.nrCols {
		m.arr = append(m.arr, row...)
		m.nrRows += 1
		return nil
	}
	return errors.New("length of appended row does not match number of columns")
}

func (m Matrix) Get(row int, col int) (int, error) {
	if m.inside(row, col) {
		return m.arr[m.index(row, col)], nil
	}
	return -1, errors.New("Get of value outside matrix boundary")
}

func (m *Matrix) Set(row int, col int, value int) error {
	if m.inside(row, col) {
		m.arr[m.index(row, col)] = value
		return nil
	}
	return errors.New("Set on point outside matrix boundry")
}

func (m Matrix) index(row int, col int) int {
	return m.nrCols*row + col
}

func (m Matrix) positionFromIndex(index int) Point {
	row := index / m.nrCols
	col := index % m.nrCols
	return Point{X: row, Y: col}
}

func (m Matrix) inside(row int, col int) bool {
	insideRows := 0 <= row && row < m.nrRows
	insideCols := 0 <= col && col < m.nrCols
	return insideRows && insideCols
}

func (m Matrix) String() string {
	out := ""
	for i := range m.arr {
		if i%(m.nrCols) == 0 {
			out += fmt.Sprintf("\n")
		}
		rep := fmt.Sprint(m.arr[i])
		if len(rep) == 1 {
			rep += " "
		}
		out += fmt.Sprintf(" : " + rep)
	}
	return out
}

func (m Matrix) FindFirst(value int) Point {
	for i := range m.arr {
		if m.arr[i] == value {
			return m.positionFromIndex(i)
		}
	}
	return Point{0, 0}
}

func (m Matrix) GetNrRows() int {
	return m.nrRows
}

func (m Matrix) GetNrCols() int {
	return m.nrCols
}
