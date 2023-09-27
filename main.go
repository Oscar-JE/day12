package main

import (
	"day12/mountain"
	"day12/parse"
	"fmt"
)

func main() {
	//m, start, end := parse.Parse("input_short.txt")
	m, start, end := parse.Parse("input.txt")
	//m, start, end := parse.Parse("input_test_medium.txt")
	//m, start, end := parse.Parse("day12cc.txt")
	fmt.Println(m)
	//fmt.Println(start)
	//fmt.Println(end)
	mountain := mountain.Init(m, start)
	mountain.CalcPathValues()
	mountain.PrintPathValues()
	fmt.Println(mountain.GetPathValue(end.X, end.Y))
}
