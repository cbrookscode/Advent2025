package Day2

import (
	"strconv"
	"strings"
)

func CalcInvalid(input []string) int {
	total := 0
	for _, numberSet := range input {
		// Get number range, split around "-", and convert to ints
		splitNums := strings.Split(numberSet, "-")
		start, _ := strconv.Atoi(splitNums[0])
		end, _ := strconv.Atoi(splitNums[1])

		// loop over number range and ignore non symetrical numbers. Make each number string, split along midpoint and compare each side.
		for i := start; i <= end; i++ {
			strNum := strconv.Itoa(i)
			if len(strNum)%2 == 0 { // is symmetrical
				mid := len(strNum) / 2
				first := strNum[0:mid]
				second := strNum[mid:]
				if first == second {
					total += i
				}
			}
		}
	}
	return total
}
