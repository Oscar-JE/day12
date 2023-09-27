package parse

import (
	"testing"
)

func TestConvertToHeight(t *testing.T) {
	res1 := []rune("abcdefghijklmnopqrstuvwxyz")
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	for index := range res1 {
		if expected[index] != ConvertToHeight(res1[index]) {
			t.Errorf("alphafan")
		}
	}
}

func TestConvert(t *testing.T) {
	in := "abc"
	expected := []int{0, 1, 2}
	out := convert(in)
	if !slizeEqual(expected, out) {
		t.Errorf("converted from string failed")
	}
}

func slizeEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	elEqual := true
	for i := range a {
		elEqual = elEqual && a[i] == b[i]
	}
	return elEqual
}

func TestParse(t *testing.T) {
	m, _, _ := Parse("../input_short.txt")
	hight := m.Get(0, 1)
	if hight != 0 {
		t.Errorf("wrong behavior in pars method")
	}
}
