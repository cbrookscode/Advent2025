package Day2

import (
	"fmt"
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

		// loop over number range and ignore non symetrical numbers. Make each number string, split along midpoint and compare each side.
		for i := start; i <= end; i++ {
			if i == 666666666 {
				fmt.Println("devil")
			}
			strNum := strconv.Itoa(i)
			if len(strNum)%2 == 0 { // is symmetrical
				mid := len(strNum) / 2
				first := strNum[0:mid]
				second := strNum[mid:]
				if first == second {
					total += i
				} else {
					splits := len(strNum) / 2
					last := strNum[0:2]
					for j := 2; j < len(strNum); j + 2 {
						if strNum[j:j+1] != last {
							symmetrical = false
						}
					}
				}
			} else if len(strNum)%3 == 0 {
				splits := len(strNum) / 3
				if splits > 3 {
					fmt.Println("greater than 3")
				}
				symmetrical := true
				last := strNum[0:splits]
				if len(strNum) == 3 {
					for j := 1; j < len(strNum); j++ {
						if strNum[j:j+1] != last {
							symmetrical = false
						}
					}
					if symmetrical {
						fmt.Println(i)
						total += i
						continue
					}
				} else {
					for j := splits; j <= len(strNum)-splits; j += splits {
						if strNum[j:j+splits] != last {
							symmetrical = false
						}
					}
					if symmetrical {
						fmt.Println(i)
						total += i
					}
					if !symmetrical {
						symmetrical = true
						last := strNum[0:3]
						for j := 3; j <= len(strNum)-3; j += 3 {
							if strNum[j:j+3] != last {
								symmetrical = false
							}
						}
						if symmetrical {
							fmt.Println(i)
							total += i
						}
					}
				}
			}
		}
	}
	return total
}

// func CalcInvalidPartTwo() int {

// }
