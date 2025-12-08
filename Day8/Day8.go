package Day8

import "math"

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
	if float64(math.Abs(point.Coord(kdNode.axis)-kdNode.loc.Coord(kdNode.axis))) < best {
		otherpotBest, otherpotBestNode := FindClosestNode(point, best, bestNode, second)
		if otherpotBest < best {
			best = otherpotBest
			bestNode = otherpotBestNode
		}
	}
	return best, bestNode
}
