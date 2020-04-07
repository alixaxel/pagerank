/*
Package pagerank implements the *weighted* PageRank algorithm.
*/
package pagerank

// Node64 is a node in a graph
type Node64 struct {
	weight   [2]float64
	outbound float64
	edges    map[uint]float64
}

// Graph64 holds node and edge data.
type Graph64 struct {
	count uint
	index map[uint64]uint
	nodes []Node64
}

// NewGraph64 initializes and returns a new graph.
func NewGraph64(size ...int) *Graph64 {
	capacity := 8
	if len(size) == 1 {
		capacity = size[0]
	}
	return &Graph64{
		index: make(map[uint64]uint, capacity),
		nodes: make([]Node64, 0, capacity),
	}
}

// Link creates a weighted edge between a source-target node pair.
// If the edge already exists, the weight is incremented.
func (g *Graph64) Link(source, target uint64, weight float64) {
	s, ok := g.index[source]
	if !ok {
		s = g.count
		g.index[source] = s
		g.nodes = append(g.nodes, Node64{})
		g.count++
	}

	g.nodes[s].outbound += weight

	t, ok := g.index[target]
	if !ok {
		t = g.count
		g.index[target] = t
		g.nodes = append(g.nodes, Node64{})
		g.count++
	}

	if g.nodes[s].edges == nil {
		g.nodes[s].edges = map[uint]float64{}
	}

	g.nodes[s].edges[t] += weight
}

// Rank computes the PageRank of every node in the directed graph.
// α (alpha) is the damping factor, usually set to 0.85.
// ε (epsilon) is the convergence criteria, usually set to a tiny value.
//
// This method will run as many iterations as needed, until the graph converges.
func (g *Graph64) Rank(α, ε float64, callback func(id uint64, rank float64)) {
	Δ := float64(1.0)
	nodes := g.nodes
	inverse := 1 / float64(len(nodes))

	// Normalize all the edge weights so that their sum amounts to 1.
	for _, node := range nodes {
		if outbound := node.outbound; outbound > 0 {
			for target := range node.edges {
				node.edges[target] /= outbound
			}
		}
	}

	leak := float64(0)

	a, b := 0, 1
	for source, node := range nodes {
		nodes[source].weight[a] = inverse

		if node.outbound == 0 {
			leak += inverse
		}
	}

	for Δ > ε {
		adjustment := (1-α)*inverse + α*leak*inverse
		for source, node := range nodes {
			aa, bb := α*node.weight[a], node.weight[b]
			for target, weight := range node.edges {
				nodes[target].weight[b] += aa * weight
			}

			nodes[source].weight[b] = bb + adjustment
		}

		Δ, leak = 0, 0
		for source, node := range nodes {
			aa, bb := node.weight[a], node.weight[b]
			if difference := aa - bb; difference < 0 {
				Δ -= difference
			} else {
				Δ += difference
			}

			if node.outbound == 0 {
				leak += bb
			}
			nodes[source].weight[a] = 0
		}

		a, b = b, a
	}

	for key, value := range g.index {
		callback(key, nodes[value].weight[a])
	}
}

// Reset clears all the current graph data.
func (g *Graph64) Reset(size ...int) {
	capacity := 8
	if len(size) == 1 {
		capacity = size[0]
	}
	g.count = 0
	g.index = make(map[uint64]uint, capacity)
	g.nodes = make([]Node64, 0, capacity)
}
