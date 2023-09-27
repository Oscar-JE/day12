package mountain

import (
	"day12/matrix"
	"testing"
)

func TestReachablePointsFromUnExplored(t *testing.T) {
	m := matrix.FilledWith(3, 3, -1)
	mount := Init(m, matrix.Point{1, 1})
	reachable := mount.reachablePointsFrom(matrix.Point{1, 1})
	if !(len(reachable) == 4) {
		t.Errorf("wrong number of reachable points")
	}
}

func TestReachablePointsFrom(t *testing.T) {
	m := matrix.FilledWith(3, 3, 10)
	mount := Init(m, matrix.Point{1, 1})
	reachable := mount.reachablePointsFrom(matrix.Point{1, 1})
	if !(len(reachable) == 4) {
		t.Errorf("wrong number of reachable points")
	}
}

func TestReachablePointsFromInHole(t *testing.T) {
	m := matrix.FilledWith(3, 3, 10)
	m.Set(1, 1, 5)
	mount := Init(m, matrix.Point{1, 1})
	reachable := mount.reachablePointsFrom(matrix.Point{1, 1})
	if !(len(reachable) == 0) {
		t.Errorf("wrong number of reachable points")
	}
}


