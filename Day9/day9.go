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
