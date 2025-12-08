package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cbrookscode/Advent2025/Day7"
)

func main() {
	file, err := os.Open("Day7/Day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := [][]string{}

	startfound := false
	startlocation := 0

	for scanner.Scan() {
		line := scanner.Text()
		if !startfound {
			for i, char := range line {
				if string(char) == "S" {
					startfound = true
					startlocation = i
				}
			}
		} else {
			listtoAdd := []string{}
			for _, char := range line {
				listtoAdd = append(listtoAdd, string(char))
			}
			input = append(input, listtoAdd)
		}
	}

	total := Day7.CountNumOfAvailPaths(startlocation, input)
	fmt.Println(total)
}
