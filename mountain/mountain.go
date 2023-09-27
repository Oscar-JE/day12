package mountain

import (
	"day12/matrix"
	"fmt"
)

type Mountain struct {
	heightMap       matrix.Matrix
	start           matrix.Point
	pathValuesToEnd matrix.Matrix
}

func Init(heightMap matrix.Matrix, end matrix.Point) Mountain {
	pathValuesToEnd := matrix.FilledWith(heightMap.GetNrRows(), heightMap.GetNrCols(), -1)
	return Mountain{heightMap: heightMap, start: end, pathValuesToEnd: pathValuesToEnd}
}

func (mont Mountain) CalculateShortestPathToPoint(start matrix.Point) int {
	mont.CalcPathValues()
	return mont.pathValuesToEnd.Get(start.X, start.Y)
}

func (mont *Mountain) CalcPathValues() {
	mont.calcPathValues(mont.start, 0)
}

func (mont *Mountain) calcPathValues(point matrix.Point, nrOfStepsToEnd int) {
	mont.pathValuesToEnd.Set(point.X, point.Y, nrOfStepsToEnd)
	reachable := mont.reachablePointsFrom(point)
	for _, loopPoint := range reachable {
		nextValue := nrOfStepsToEnd + 1
		pathValueAtCandidate := mont.pathValuesToEnd.Get(loopPoint.X, loopPoint.Y)
		if pathValueAtCandidate == -1 || nextValue < pathValueAtCandidate {
			mont.calcPathValues(loopPoint, nextValue)
		}
	}
}

func (mont Mountain) reachablePointsFrom(point matrix.Point) []matrix.Point {
	var diff []matrix.Point = []matrix.Point{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}}
	reachablePoints := []matrix.Point{}
	heightAtPoint := mont.heightMap.Get(point.X, point.Y)
	for _, d := range diff {
		candidate := point.Add(d)
		if mont.heightMap.InMatrix(candidate) {
			hightAtCandidate := mont.heightMap.Get(candidate.X, candidate.Y)
			heightDiff := hightAtCandidate - heightAtPoint
			if validPoint(heightDiff) {
				reachablePoints = append(reachablePoints, candidate)
			}
		}

	}
	return reachablePoints
}

func validPoint(heightDiff int) bool {
	return heightDiff <= 1
}

func (mont Mountain) PrintPathValues() {
	fmt.Println(mont.pathValuesToEnd)
}

func (mont Mountain) GetPathValue(row int, col int) int {
	value := mont.pathValuesToEnd.Get(row, col)
	return value
}
