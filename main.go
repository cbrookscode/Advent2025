package main

import (
	"fmt"

	"github.com/cbrookscode/Advent2025/Day8"
)

func main() {
	kdtree := Day8.BuildKBTree(0, nil, Day8.Testcase)
	dist, closPoint := Day8.FindClosestNode(Day8.Point{X: 162, Y: 817, Z: 812}, 90000000000000, Day8.Point{}, kdtree)
	fmt.Println(closPoint)
	fmt.Println(dist)
}
