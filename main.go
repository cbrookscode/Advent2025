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

	pairs, strmap := Day8.GetPairs(input)
	_, nums := Day8.GetCircuits(pairs, strmap)
	fmt.Println(nums)
	fmt.Println(nums[0] * nums[1])
}
