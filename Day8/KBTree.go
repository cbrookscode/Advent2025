package Day8

import (
	"sort"
)

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
	axis := depth % 3

	// base case to stop recusion
	if len(points) == 0 {
		return nil
	}

	sort.Slice(points, func(i, j int) bool { return int(points[i].Coord(axis)) < int(points[j].Coord(axis)) })

	mid := len(points) / 2
	left := points[:mid]
	right := []Point{}
	if mid+1 < len(points) {
		right = points[mid+1:]
	}

	newNode := &Node{
		parent: parent,
		loc:    points[mid],
	}
	newNode.left = BuildKBTree(depth+1, newNode, left)
	newNode.right = BuildKBTree(depth+1, newNode, right)
	return newNode
}
