package main

import (
	"day12/parse"
	"fmt"
)

func main() {
	m, start, end := parse.Parse("input_short.txt")
	fmt.Println(m)
	fmt.Println(start)
	fmt.Println(end)
}
