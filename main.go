package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cbrookscode/Advent2025/Day8"
)

func main() {
	file, err := os.Open("Day8/Day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	input := []Day8.Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		input = append(input, Day8.Point{X: x, Y: y, Z: z})
	}

	// kdtree := Day8.BuildKBTree(0, nil, input)
	pairs, strmap := Day8.GetPairs(input)
	circuits := Day8.GetCircuits(pairs)
	fmt.Println(len(circuits))
	first := strmap[circuits[0][0]][0]
	sec := strmap[circuits[0][len(circuits[0])-1]][0]
	fmt.Println(first * sec)

	// no need for kd tree impelementation. need distance values for all possible point combinations and then sort that list and feed into union find for solution to part 1.
}
