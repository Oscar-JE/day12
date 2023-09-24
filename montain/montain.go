package montain

import "day12/matrix"

type Montain struct {
	heightMap       matrix.Matrix
	end             matrix.Point
	pathValuesToEnd matrix.Matrix
}

func Init(heightMap matrix.Matrix, end matrix.Point) Montain {
	return Montain{heightMap: heightMap, end: end}
}

func (mont Montain) CalculateShortestPathToPoint(start matrix.Point) (int, error) {
	mont.CalcPathValues()
	return mont.pathValuesToEnd.Get(start.X, start.Y)
}

func (mont *Montain) CalcPathValues() {
	//mont.calcPathValues(mont.end, 0)
}

func reachablePointsFrom(point matrix.Point) []matrix.Point {

}
