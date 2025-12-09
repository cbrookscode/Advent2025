package Day8

import (
	"math"
	"sort"
	"strconv"
)

type Point struct {
	X int
	Y int
	Z int
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

func (p *Point) Stringify() string {
	return strconv.Itoa(p.X) + strconv.Itoa(p.Y) + strconv.Itoa(p.Z)
}

type Node struct {
	loc   Point
	axis  int
	left  *Node
	right *Node
}

func FindClosestNode(point Point, curr_best float64, best_node Point, kdNode *Node) (float64, Point) {
	if kdNode == nil {
		return curr_best, best_node
	}

	best := curr_best
	bestNode := best_node
	dx := float64((kdNode.loc.X - point.X) * (kdNode.loc.X - point.X))
	dy := float64((kdNode.loc.Y - point.Y) * (kdNode.loc.Y - point.Y))
	dz := float64((kdNode.loc.Z - point.Z) * (kdNode.loc.Z - point.Z))
	newDist := math.Sqrt(dx + dy + dz)
	if newDist != 0 && newDist < curr_best {
		best = newDist
		bestNode = kdNode.loc
	}

	first := &Node{}
	second := &Node{}
	if point.Coord(kdNode.axis) < kdNode.loc.Coord(kdNode.axis) {
		first = kdNode.left
		second = kdNode.right
	} else {
		first = kdNode.right
		second = kdNode.left
	}

	potBest, potBestNode := FindClosestNode(point, best, bestNode, first)
	if potBest < best {
		best = potBest
		bestNode = potBestNode
	}

	// if distance on axis to splitting plane is less then "radius" best closest found so far, then the excess is bleeding over to other side and need to check it
	if float64(math.Abs(point.Coord(kdNode.axis)-kdNode.loc.Coord(kdNode.axis))) < best {
		otherpotBest, otherpotBestNode := FindClosestNode(point, best, bestNode, second)
		if otherpotBest < best {
			best = otherpotBest
			bestNode = otherpotBestNode
		}
	}
	return best, bestNode
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
		loc: points[mid],
	}
	newNode.left = BuildKBTree(depth+1, newNode, left)
	newNode.right = BuildKBTree(depth+1, newNode, right)
	return newNode
}
