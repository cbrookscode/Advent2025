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

func BiggerNum(input [][]int) int {
	total := 0
	bankLen := 12
	for _, numbers := range input {
		final := []int{}
		interval := len(numbers) - bankLen
		remaining := interval
		for j := 0; j < len(numbers); j++ {
			if len(final) == bankLen {
				break
			}
			index := 0
			if j+interval > len(numbers)-1 {
				index = getNums(numbers[j:], remaining)
			} else {
				index = getNums(numbers[j:j+interval], remaining)
			}

			if j+index+1 == len(numbers)-1 {
				fmt.Println(numbers[j:])
				final = append(final, numbers[index-remaining:index+1]...) // struggling to figure out how to append based on the numbers being skipped and amount of skips left
				remaining -= index
				continue

			}
			fmt.Println(numbers[j : j+index+1])
			final = append(final, numbers[j:j+index+1]...)
			j += index
			remaining -= index
		}
		toAdd := TotalUp(final)
		fmt.Println(toAdd)
		total += toAdd
	}
	return total
}

// need to know remaining skips, because if not enough skips left then need to chose next best number
func getNums(nums []int, remainingSkips int) int {
	biggest := nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		if i > remainingSkips {
			return index
		}
		if nums[i] > biggest {
			biggest = nums[i]
			index = i
		}
	}
	if index == len(nums)-1 {
		return index
	}
	return index
}
func TotalUp(nums []int) int {
	total := nums[0]
	for j := 1; j < len(nums); j++ {
		total = (total * 10) + nums[j]
	}
	return total
}
