package Day8

import "sort"

type Pair struct {
	dist float64
	p1   string
	p2   string
}

func GetShortestDistances(points []Point, kdtree *Node) []Pair {
	pairs := []Pair{}

	for _, point := range points {
		dist, closest := FindClosestNode(point, 90000000000000000, Point{}, kdtree)
		p1 := point.Stringify()
		p2 := closest.Stringify()
		pairExists := false
		for _, pair := range pairs {
			if (pair.p1 == p1 && pair.p2 == p2) || (pair.p1 == p2 && pair.p2 == p1) {
				// already made this pair
				pairExists = true
			}
		}
		if !pairExists {
			pairs = append(pairs, Pair{dist: dist, p1: p1, p2: p2})
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].dist < pairs[j].dist })
	return pairs
}

func GetCircuits(orderedPairs []Pair) [][]string {
	uf := BuildUnionFind()
	for i, pair := range orderedPairs {
		if i == 10 {
			break
		}
		uf.Union(pair.p1, pair.p2)
	}
	return uf.Group()
}

func CircuitTotal(listOfCircuits [][]string) int {
	first := 0
	second := 0
	third := 0
	for _, cirList := range listOfCircuits {
		cirTotal := len(cirList)
		if cirTotal > first {
			third = second
			second = first
			first = cirTotal
		} else if cirTotal > second {
			third = second
			second = cirTotal
		} else if cirTotal > third {
			third = cirTotal
		}
	}
	return first * second * third
}
