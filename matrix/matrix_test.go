package matrix

import (
	"testing"
)

func TestFilledWith(t *testing.T) {
	m := FilledWith(2, 5, 8)
	m.Set(0, 0, 3)
	if m.Get(0, 1) != 8 {
		t.Errorf("All values changed incorrectly")
	}
}

func TestAppend(t *testing.T) {
	m := Matrix{[][]int{}}
	m.Append([]int{1, 2, 3})
	if m.Get(0,0) != 1 && m.Get(0,2) != 3{
		t.Errorf("append failed")
	}
	m.Append([]int{4,5,6})
	if m.Get(0,0) != 1 && m.Get(0,2) != 3{
		t.Errorf("append failed")
	}
	if m.Get(1,0) != 4 && m.Get(1,2) != 6{
		t.Errorf("append failed")
	}
}

func TestFindFirst(t *testing.T){
	m := Matrix{[][]int{}}
	m.Append([]int{1, 2, 3})
	m.Append([]int{4,5,6})
	res1 := m.FindFirst(1)
	res2 := m.FindFirst(5)
	if res1.X != 0 || res1.Y != 0 || res2.X != 1 || res2.Y != 1{
		t.Errorf("error in find first")
	} 
}