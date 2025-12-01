package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := countTimeHitZero(test1)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("final count: %v\n", num)
}

func countTimeHitZero(rotationList []string) (int, error) {
	startPoint := 50
	count := 0
	for _, rotation := range rotationList {
		direction := rotation[0:1]
		num, err := strconv.ParseInt(rotation[1:], 10, 32)
		if err != nil {
			return 0, err
		}

		constrainedNum := int(num) % 100
		count += int(num) / 100
		if constrainedNum == 0 {
			continue
		}

		if startPoint == 0 {
			count++
			startPoint, _ = CalcPosition(startPoint, constrainedNum, count, direction)
			continue
		}

		startPoint, count = CalcPosition(startPoint, constrainedNum, count, direction)
	}
	return count, nil
}

func CalcPosition(current, change, count int, direction string) (newLocation int, newCount int) {
	switch direction {
	case "L":
		if current-change < 0 {
			if current != 0 {
				count++
			}
			current = 100 - (change - current)
		} else {
			current = current - change
		}
	case "R":
		if current+change > 99 {
			current = (change - (100 - current))
			if current != 0 {
				count++
			}
		} else {
			current = current + change
		}
	}
	return current, count
}
