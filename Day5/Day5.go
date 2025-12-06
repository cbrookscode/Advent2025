package Day5

import (
	"fmt"
	"sort"
)

func GirthMasterFormula(input [][]int64) [][]int64 {
	sort.Slice(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})

	merged := [][]int64{}
	current := []int64{input[0][0], input[0][1]}

	for i := 1; i < len(input); i++ {
		nextLow := input[i][0]
		nextHigh := input[i][1]

		if nextLow <= current[1] { //overlap garunteed because already sorted
			if nextLow < current[0] {
				current[0] = nextLow
			}

			if nextHigh > current[1] {
				current[1] = nextHigh
			}
		} else { // no overlap
			merged = append(merged, current)
			current = []int64{nextLow, nextHigh}
		}
	}

	merged = append(merged, current)
	return merged
}

func GetAllNumbersInProvidedRanges(input [][]int64) int64 {
	fmt.Println(input)
	var total int64

	for _, minMax := range input {
		max := minMax[1]
		min := minMax[0]
		total += (max - min) + 1
	}
	return total
}
