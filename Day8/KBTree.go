package Day8

import "sort"

type Point struct {
	X int
	Y int
	Z int
}
type Node struct {
	loc    Point
	axis   int
	parent *Node
	left   *Node
	right  *Node
}

func (p *Point) Coord(i int) float64 {
	switch i {
	case 0:
		return float64(p.X)
	case 1:
		return float64(p.Y)
	case 2:
		return float64(p.Z)
	default:
		panic("invalid aXis")
	}
}

func BuildKBTree(depth int, parent *Node, points []Point) *Node {
	axis := depth % 2

	// base case to stop recusion
	if len(points) == 1 {
		return &Node{
			loc:    points[0],
			axis:   axis,
			parent: parent,
			left:   nil,
			right:  nil,
		}
	} else if points == nil {
		return nil
	}

	if axis == 0 { // split along x
		// sort list by x
		sort.Slice(points, func(i, j int) bool { return points[i].X < points[j].X })
	}

	if axis == 1 { // split along y
		// sort list by y
		sort.Slice(points, func(i, j int) bool { return points[i].Y < points[j].Y })
	}

	if axis == 2 { // split along z
		// sort list by z
		sort.Slice(points, func(i, j int) bool { return points[i].Z < points[j].Z })
	}

	mid := len(points) / 2
	left := points[:mid]
	right := points[mid+1:]

	newNode := &Node{
		parent: parent,
		loc:    points[mid],
	}
	newNode.left = BuildKBTree(depth+1, newNode, left)
	newNode.right = BuildKBTree(depth+1, newNode, right)
	return newNode
}
