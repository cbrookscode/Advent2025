package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cbrookscode/Advent2025/Day5"
)

func main() {
	file, err := os.Open("Day5/Day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := [][]int64{}
	ingredientNums := []int64{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			num1, _ := strconv.ParseInt(parts[0], 10, 64)
			num2, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, []int64{num1, num2})

		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			ingredientNums = append(ingredientNums, num)
		}
	}
	fmt.Println("=====================================")
	minMaxlist := Day5.GirthMasterFormula(ranges)
	fmt.Println(Day5.GetAllNumbersInProvidedRanges(minMaxlist))

}
