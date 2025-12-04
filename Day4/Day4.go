package Day4

func TotalClearedNodes(grid [][]string) int {
	total := 0
	zeroRemaining := false
	for !zeroRemaining {
		toBeRemoved := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "@" {
					paperRollTotal := CheckSurroundingNodes(i, j, grid)
					if paperRollTotal < 4 {
						grid[i][j] = "p"
						toBeRemoved++
					}
				}
				if grid[i][j] == "p" {
					grid[i][j] = "x"
					total++
				}
			}
		}
		if toBeRemoved == 0 {
			zeroRemaining = true
		}
	}
	return total
}

func CheckSurroundingNodes(i, j int, grid [][]string) int {
	total := 0
	if i-1 >= 0 {
		if j-1 >= 0 {
			TL := grid[i-1][j-1]
			if TL == "@" {
				total++
			}
		}
		if j+1 < len(grid[i]) {
			TR := grid[i-1][j+1]
			if TR == "@" {
				total++
			}
		}
		T := grid[i-1][j]
		if T == "@" {
			total++
		}
	}

	if j-1 >= 0 {
		L := grid[i][j-1]
		if L == "@" {
			total++
		}
		if i+1 < len(grid) {
			BL := grid[i+1][j-1]
			if BL == "@" {
				total++
			}
		}
	}
	if j+1 < len(grid[i]) {
		R := grid[i][j+1]
		if R == "@" {
			total++
		}
		if i+1 < len(grid) {
			BR := grid[i+1][j+1]
			if BR == "@" {
				total++
			}
		}
	}

	if i+1 < len(grid) {
		B := grid[i+1][j]
		if B == "@" {
			total++
		}
	}
	return total
}
