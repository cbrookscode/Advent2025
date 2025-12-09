package main

import (
	"fmt"

	"github.com/cbrookscode/Advent2025/Day8"
)

func main() {
	kdtree := Day8.BuildKBTree(0, nil, Day8.Testcase)
	pairs := Day8.GetShortestDistances(Day8.Testcase, kdtree)
	// fmt.Println(pairs)
	circuits := Day8.GetCircuits(pairs)
	fmt.Println(circuits)
	total := Day8.CircuitTotal(circuits)
	fmt.Println(total)
}
