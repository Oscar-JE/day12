package matrix

import (
	"testing"
)

func TestWrongAppend(t *testing.T) {
	m := Init(2)
	row := []int{1, 2, 3}
	if m.Append(row) == nil {
		t.Errorf("no error raised when preforming illegal append")
	}
}
