package main

import (
	"day12/montain"
	"day12/parse"
	"fmt"
)

func main() {
	//m, start, end := parse.Parse("input_short.txt")
	m, start, end := parse.Parse("input.txt")
	fmt.Println(m)
	fmt.Println(start)
	fmt.Println(end)
	montain := montain.Init(m, end)
	montain.PrintPathValues()
	montain.CalcPathValues()
	montain.PrintPathValues()
	fmt.Println(montain.GetPathValue(start.X, start.Y))
}
