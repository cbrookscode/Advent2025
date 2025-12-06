package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cbrookscode/Advent2025/Day6"
)

func main() {
	file, err := os.Open("Day6/Day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := map[int][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			input[i] = append(input[i], string(char))
		}
	}
	total := Day6.CalcLineAndTotalUp(input)
	fmt.Println(total)
}
