package Day2

import (
	"strconv"
	"strings"
)

// Part One Solution
func CalcInvalid(input []string) int {
	total := 0
	for _, numberSet := range input {
		// Get number range, split around "-", and convert to ints
		splitNums := strings.Split(numberSet, "-")
		start, _ := strconv.Atoi(splitNums[0])
		end, _ := strconv.Atoi(splitNums[1])

		for i := start; i <= end; i++ {
			// single digits are valid
			if i < 11 {
				continue
			}
			strNum := strconv.Itoa(i)
			if isNumberInvalid(strNum) {
				total += i
			}
		}
	}
	return total
}

/* check applicable increments for a given string of numbers. List of numbers is len(s)/2 and result has no remainder when modulo'd against current value of i in loop
 and then see if increment yields an invalid number */
func isNumberInvalid(s string) bool {
	n := len(s)
	// Loop over potential viable increments
	for i := 1; i <= n/2; i++ {
		if n % i != 0 { // only keep those that have no remainder against current i value
			continue
		}
		repeat := true
		for j := i; j < n; j += i { // test to see if i value yields a repeated string across its full length
			if s[j:j+i] != s[:i] {
				repeat = false
				break
			}
		}
		if repeat {
			return true
		}
	}
	return false
}
