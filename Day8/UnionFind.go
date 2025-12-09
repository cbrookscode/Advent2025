package Day8

type UnionFind struct {
	parent map[string]string
	rank   map[string]int
}

func BuildUnionFind() *UnionFind {
	return &UnionFind{parent: make(map[string]string), rank: make(map[string]int)}
}

func (u *UnionFind) Find(x string) (parent string) {
	// if doesn't exists, create in UF struct and return input
	if _, exists := u.parent[x]; !exists {
		u.parent[x] = x
		u.rank[x] = 0
		return x
	}

	/* if we arent finding on a root node, call find on the parent of the passed x. This will recusively pass down the actual
	parent of this group to each child directly in the chain so that future lookups are O(1) */
	if u.parent[x] != x { // if x is not a parent
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFind) Union(x, y string) {
	// need the parents of each, merge at roots, type of merge is dependent on rank
	rootX := u.Find(x)
	rootY := u.Find(y)

	if rootX == rootY { // already in a group together
		return
	}

	if u.rank[rootX] < u.rank[rootY] {
		u.parent[rootX] = rootY
	} else if u.rank[rootX] > u.rank[rootY] {
		u.parent[rootY] = rootX
	} else { // only need to increase rank when equal since rank is relative and not abs height. only increase rank of the new parent.
		u.parent[rootX] = rootY
		u.rank[rootY]++
	}
}

func (u *UnionFind) Group() [][]string {
	groups := make(map[string][]string)

	for node := range u.parent {
		root := u.Find(node)
		// for every node, find its root and add the node to that roots list of strings
		groups[root] = append(groups[root], node)
	}

	listOfGroups := make([][]string, 0, len(groups))
	for _, children := range groups {
		listOfGroups = append(listOfGroups, children)
	}
	return listOfGroups
}
