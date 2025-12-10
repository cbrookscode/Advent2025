package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cbrookscode/Advent2025/Day9"
)

func main() {
	file, err := os.Open("Day9/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	input := []Day9.Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		input = append(input, Day9.Point{X: x, Y: y})
	}
	area := Day9.GirMagic(input)
	fmt.Println(area)
}
