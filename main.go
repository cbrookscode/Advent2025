package main

import (
	"fmt"

	"github.com/cbrookscode/Advent2025/Day8"
)

func main() {
	// file, err := os.Open("Day6/Day6.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// input := map[int][]string{}
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	for i, char := range line {
	// 		input[i] = append(input[i], string(char))
	// 	}
	// }
	// total := Day6.CalcLineAndTotalUp(input)
	// fmt.Println(total)
	kdtree := Day8.BuildKBTree(0, nil, Day8.Testcase)
	value, point := Day8.FindClosestNode(Day8.Point{X: 57, Y: 618, Z: 57}, 9000000000000000000000000, Day8.Point{}, kdtree)
	fmt.Println(point)
	fmt.Println(value)
}
