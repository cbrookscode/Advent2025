package Day1

import (
	"strconv"
)

func countTimePassZero(rotationList []string) (int, error) {
	pos := 50
	count := 0
	for _, rotation := range rotationList {
		direction := rotation[0:1]
		dist, err := strconv.ParseInt(rotation[1:], 10, 32)
		if err != nil {
			return 0, err
		}

		count += int(dist) / 100
		dist = dist % 100
		old := pos

		if direction == "L" {
			pos -= int(dist)
		} else {
			pos += int(dist)
		}

		if old != 0 && pos > 100 {
			pos = pos % 100
			count++
		} else if old != 0 && pos < 0 {
			pos = (pos % 100) * -1
			count++
		}
		if pos == 0 {
			count++
		}
	}
	return count, nil
}
