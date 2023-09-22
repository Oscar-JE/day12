package matrix

import (
	"errors"
)

type Matrix struct {
	arr    []int
	nrRows int
	nrCols int
}

func Init(nrCols int) Matrix {
	return Matrix{[]int{}, 0, nrCols}
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

func (m Matrix) inside(row int, col int) bool {
	insideRows := 0 <= row && row < m.nrRows
	insideCols := 0 <= col && col < m.nrCols
	return insideRows && insideCols
}
