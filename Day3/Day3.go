package Day3

import "fmt"

func BigNum(input [][]int) int {
	total := 0
	for _, numbers := range input {
		leftDigit := numbers[0]
		rightDigit := 0
		for j := 1; j < len(numbers); j++ {
			if numbers[j] > leftDigit && j != len(numbers)-1 {
				leftDigit = numbers[j]
				rightDigit = numbers[j+1]
			} else if numbers[j] > rightDigit {
				rightDigit = numbers[j]
			}
		}
		fmt.Printf("%v\n", (leftDigit*10)+rightDigit)
		total += (leftDigit * 10) + rightDigit
	}
	return total
}

func TotalUp(nums []int) int {
	total := nums[0]
	for j := 1; j < len(nums); j++ {
		total = (total * 10) + nums[j]
	}
	return total
}

func Dumbass(input [][]int) int {
	total := 0
	bankLen := 12
	for _, numbers := range input {
		final := []int{}
		left := 0
		right := len(numbers) - bankLen

		for len(final) < bankLen {
			pos := left
			maxVal := 0
			maxValPos := 0
			for pos <= right {
				if numbers[pos] > maxVal {
					maxVal = numbers[pos]
					maxValPos = pos
				}
				pos++
			}
			final = append(final, maxVal)
			left = maxValPos + 1
			right++
		}
		total += TotalUp(final)
	}
	return total
}
