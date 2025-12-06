package Day6

import (
	"strconv"
)

func CalcLineAndTotalUp(input map[int][]string) int {
	total := 0
	numList := []int{}
	for i := len(input) - 1; i >= 0; i-- {
		numsToCombine := []string{}
		reset := false
		for _, char := range input[i] {
			numsToCombine = append(numsToCombine, string(char))
			if char == "*" || char == "+" {
				numList = append(numList, combineNum(numsToCombine))
				total += domath(numList, char)
				numList = []int{}
				reset = true
			}
		}
		if !reset {
			numList = append(numList, combineNum(numsToCombine))
		}
	}
	return total
}

func domath(nums []int, oper string) int {
	if oper == "*" {
		total := 1
		for _, num := range nums {
			if num == 0 {
				continue
			}
			total *= num
		}
		return total
	}

	if oper == "+" {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	return 0
}

func combineNum(nums []string) int {
	total := 0
	for _, strNum := range nums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue
		}
		total = (total * 10) + num
	}
	return total
}
