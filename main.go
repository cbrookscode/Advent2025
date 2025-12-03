package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cbrookscode/Advent2025/Day3"
)

func main() {
	file, err := os.Open("Day3/Day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputList := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		listofNums := []int{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
				return
			}
			listofNums = append(listofNums, num)
		}

		inputList = append(inputList, listofNums)
	}
	fmt.Println(Day3.BiggerNum(Day3.TestPart2))
}
