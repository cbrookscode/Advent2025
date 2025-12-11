package Day9

type Point struct {
	X int
	Y int
}

func CalcArea(p1, p2 Point) int {
	y := (p1.Y - p2.Y)
	if y < 0 {
		y = y * -1
	}
	x := (p1.X - p2.X)
	if x < 0 {
		x = x * -1
	}
	area := (x + 1) * (y + 1)
	if area < 0 {
		return area * -1
	}
	return area
}

func GirMagic(points []Point) (area int) {
	bestAreaFound := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if points[j].X > points[i].X && points[j].Y >= points[i].Y { // possible best candidate for area
				newArea := CalcArea(points[i], points[j])
				if newArea > bestAreaFound {
					bestAreaFound = newArea
				}
			}
		}
	}
	return bestAreaFound
}

func Validate(point Point, x1, x2, y1, y2 int) {
	// if (point.x is between/equal x1 and x2) and (point.y is between/equal y1 and y2)
}

func GirsMagicPartTwo(biggestLine int, rPoints []Point, allPoints map[int]map[int]struct{}) {
	biggest := biggestLine
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if points[j].X > points[i].X && points[j].Y >= points[i].Y { // possible best candidate for area
				// validate, till bound
			}
		}
	}
	return
}

func MakePoints(p1, p2 Point, pointMap map[int]map[int]struct{}) int {
	biggest := 0
	if innerMap, exists := pointMap[p1.X]; exists {
		innerMap[p1.Y] = struct{}{}
	} else {
		inner := map[int]struct{}{}
		inner[p1.Y] = struct{}{}
		pointMap[p1.X] = inner
	}
	if innerMap, exists := pointMap[p2.X]; exists {
		innerMap[p2.Y] = struct{}{}
	} else {
		inner := map[int]struct{}{}
		inner[p2.Y] = struct{}{}
		pointMap[p2.X] = inner
	}
	if p1.X == p2.X {
		yStart := 0
		if p2.Y > p1.Y {
			yStart = p1.Y
		} else {
			yStart = p2.Y
		}
		for j := 1; j < AbsValue((p1.Y - p2.Y)); j++ {
			newP := Point{X: p2.X, Y: yStart + j}
			if innerMap, exists := pointMap[newP.X]; exists {
				innerMap[newP.Y] = struct{}{}
			} else {
				inner := map[int]struct{}{}
				inner[newP.Y] = struct{}{}
				pointMap[newP.X] = inner
			}
		}
		num := AbsValue(p1.Y-p2.Y) + 1
		if num > biggest {
			biggest = num
		}
	} else if p1.Y == p2.Y {
		xStart := 0
		if p2.X > p1.X {
			xStart = p1.X
		} else {
			xStart = p2.X
		}
		for j := 1; j < AbsValue((p1.X - p2.X)); j++ {
			newP := Point{X: xStart + j, Y: p2.Y}
			if innerMap, exists := pointMap[newP.X]; exists {
				innerMap[newP.Y] = struct{}{}
			} else {
				inner := make(map[int]struct{})
				inner[newP.Y] = struct{}{}
				pointMap[newP.X] = inner
			}
		}
		num := AbsValue(p1.X-p2.X) + 1
		if num > biggest {
			biggest = num
		}
	}
	return biggest
}

func BuildListofPoints(rPoints []Point) (map[int]map[int]struct{}, int) {
	gPoints := make(map[int]map[int]struct{})
	biggest := 0
	for i := 0; i < len(rPoints); i++ {
		if i != 0 {
			num := MakePoints(rPoints[i-1], rPoints[i], gPoints)
			if num > biggest {
				biggest = num
			}
		}
		if i+1 < len(rPoints) {
			num := MakePoints(rPoints[i+1], rPoints[i], gPoints)
			if num > biggest {
				biggest = num
			}
		}
	}
	return gPoints, biggest
}

func AbsValue(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// list of nodes provided are red, red nodes connect to each one direct before or after it but the dots connecting them are green tiles.
// the red tile list nodes that are adjacent in the list to each other (before and after) will either be same row or same column.
// the first red tile does connect to the last so that theres a loop that will then define the area within it.
// all tiles within the loop will be green tiles as well
// create a rectangle with red tiles in opp corners again but now the recntangle must all be green tiles

// horizontal or vertical lines will always be valid areas
// its when checking points that dont have the same x AND y that you have to worry if it escapes the bounds set by the loop.
// dfs when checking and if hit a none green/red point while travelling then this combo isn't valid. otherwise its good and see if its area is bigger then biggest found so far
