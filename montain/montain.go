package montain

import (
	"day12/matrix"
	"fmt"
)

type Montain struct {
	heightMap       matrix.Matrix
	end             matrix.Point
	pathValuesToEnd matrix.Matrix
}

func Init(heightMap matrix.Matrix, end matrix.Point) Montain {
	pathValuesToEnd := matrix.FilledWith(heightMap.GetNrRows(), heightMap.GetNrCols(), -1)
	return Montain{heightMap: heightMap, end: end, pathValuesToEnd: pathValuesToEnd}
}

func (mont Montain) CalculateShortestPathToPoint(start matrix.Point) (int, error) {
	mont.CalcPathValues()
	return mont.pathValuesToEnd.Get(start.X, start.Y)
}

func (mont *Montain) CalcPathValues() {
	mont.calcPathValues(mont.end, 0)
}

func (mont *Montain) calcPathValues(point matrix.Point, value int) {
	mont.pathValuesToEnd.Set(point.X, point.Y, value)
	reachable := mont.reachablePointsFrom(point, value+1)
	for _, p := range reachable {
		mont.calcPathValues(p, value+1)
	}
}

func (mont Montain) reachablePointsFrom(point matrix.Point, nextValue int) []matrix.Point {
	var diff []matrix.Point = []matrix.Point{{X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 1}}
	reachablePoints := []matrix.Point{}
	heightAtPoint, _ := mont.heightMap.Get(point.X, point.Y)
	for _, d := range diff {
		candidate := point.Add(d)
		hightAtCandidate, errHeight := mont.heightMap.Get(candidate.X, candidate.Y)
		heightDiff := hightAtCandidate - heightAtPoint
		pathValueAtCandidate, errorPath := mont.pathValuesToEnd.Get(candidate.X, candidate.Y)
		if errHeight == nil && errorPath == nil && -1 <= heightDiff && heightDiff <= 1 && (pathValueAtCandidate == -1 || nextValue < pathValueAtCandidate) {
			reachablePoints = append(reachablePoints, candidate)
		}
	}
	return reachablePoints
}

func (mont Montain) PrintPathValues() {
	fmt.Println(mont.pathValuesToEnd)
}

func (mont Montain) GetPathValue(row int, col int) int {
	value, _ := mont.pathValuesToEnd.Get(row, col)
	return value
}
