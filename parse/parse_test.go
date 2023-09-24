package parse

import (
	"testing"
)

func TestConvertToHeight(t *testing.T) {
	res1 := ConvertToHeight('a')
	res2 := ConvertToHeight('b')
	if res1 != 0 && res2 != 1 {
		t.Errorf("something wrong with convert to height")
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
	hight, err := m.Get(0, 1)
	if err != nil || hight != 0 {
		t.Errorf("wrong behavior in pars method")
	}
}
