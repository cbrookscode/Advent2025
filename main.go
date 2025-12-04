package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cbrookscode/Advent2025/Day4"
)

func main() {
	file, err := os.Open("Day4/Day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputList := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		listofNums := []string{}
		for _, char := range line {
			listofNums = append(listofNums, string(char))
		}

		inputList = append(inputList, listofNums)
	}
	fmt.Println(Day4.TotalClearedNodes(inputList))

}
