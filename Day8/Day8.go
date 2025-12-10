package Day8

import (
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	dist float64
	p1   string
	p2   string
}

func GetPairs(points []Point) []Pair {
	listofpairs := []Pair{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := float64((points[i].X - points[j].X) * (points[i].X - points[j].X))
			dy := float64((points[i].Y - points[j].Y) * (points[i].Y - points[j].Y))
			dz := float64((points[i].Z - points[j].Z) * (points[i].Z - points[j].Z))
			newDist := math.Sqrt(dx + dy + dz)
			listofpairs = append(listofpairs, Pair{dist: newDist, p1: points[i].Stringify(), p2: points[j].Stringify()})
		}
	}
	sort.Slice(listofpairs, func(i, j int) bool { return listofpairs[i].dist < listofpairs[j].dist })
	return listofpairs
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
		if i == 1000 {
			break
		}
		fmt.Printf("p1 = %v\n p2 = %v\n", pair.p1, pair.p2)
		uf.Union(pair.p1, pair.p2)
		fmt.Println(uf.Group())
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
