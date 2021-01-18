/*
Package pagerank implements the *weighted* PageRank algorithm.
*/
package pagerank

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// NumCPU is the number of cpus
	NumCPU = runtime.NumCPU()
)

// Node32 is a node in a graph
type Node32 struct {
	sync.Mutex
	weight   [2]float32
	outbound float32
	edges    map[uint]float32
}

// Graph32 holds node and edge data.
type Graph32 struct {
	Verbose bool
	count   uint
	index   map[uint64]uint
	nodes   []Node32
}

// NewGraph32 initializes and returns a new graph.
func NewGraph32(size ...int) *Graph32 {
	capacity := 8
	if len(size) == 1 {
		capacity = size[0]
	}
	return &Graph32{
		index: make(map[uint64]uint, capacity),
		nodes: make([]Node32, 0, capacity),
	}
}

// Link creates a weighted edge between a source-target node pair.
// If the edge already exists, the weight is incremented.
func (g *Graph32) Link(source, target uint64, weight float32) {
	s, ok := g.index[source]
	if !ok {
		s = g.count
		g.index[source] = s
		g.nodes = append(g.nodes, Node32{})
		g.count++
	}

	g.nodes[s].outbound += weight

	t, ok := g.index[target]
	if !ok {
		t = g.count
		g.index[target] = t
		g.nodes = append(g.nodes, Node32{})
		g.count++
	}

	if g.nodes[s].edges == nil {
		g.nodes[s].edges = map[uint]float32{}
	}

	g.nodes[s].edges[t] += weight
}

// Rank computes the PageRank of every node in the directed graph.
// α (alpha) is the damping factor, usually set to 0.85.
// ε (epsilon) is the convergence criteria, usually set to a tiny value.
//
// This method will run as many iterations as needed, until the graph converges.
func (g *Graph32) Rank(α, ε float32, callback func(id uint64, rank float32)) {
	Δ := float32(1.0)
	nodes := g.nodes
	inverse := 1 / float32(len(nodes))

	// Normalize all the edge weights so that their sum amounts to 1.
	if g.Verbose {
		fmt.Println("normalize...")
	}
	done := make(chan bool, 8)
	normalize := func(node *Node32) {
		if outbound := node.outbound; outbound > 0 {
			for target := range node.edges {
				node.edges[target] /= outbound
			}
		}
		done <- true
	}
	i, flight := 0, 0
	for i < len(nodes) && flight < NumCPU {
		go normalize(&nodes[i])
		flight++
		i++
	}
	for i < len(nodes) {
		<-done
		flight--
		go normalize(&nodes[i])
		flight++
		i++
	}
	for j := 0; j < flight; j++ {
		<-done
	}

	if g.Verbose {
		fmt.Println("initialize...")
	}
	leak := float32(0)

	a, b := 0, 1
	for source := range nodes {
		nodes[source].weight[a] = inverse

		if nodes[source].outbound == 0 {
			leak += inverse
		}
	}

	update := func(adjustment float32, node *Node32) {
		node.Lock()
		aa, bb := α*node.weight[a], node.weight[b]
		node.Unlock()
		for target, weight := range node.edges {
			nodes[target].Lock()
			nodes[target].weight[b] += aa * weight
			nodes[target].Unlock()
		}
		node.Lock()
		node.weight[b] = bb + adjustment
		node.Unlock()
		done <- true
	}
	for Δ > ε {
		if g.Verbose {
			fmt.Println("updating...")
		}
		adjustment := (1-α)*inverse + α*leak*inverse
		i, flight := 0, 0
		for i < len(nodes) && flight < NumCPU {
			go update(adjustment, &nodes[i])
			flight++
			i++
		}
		for i < len(nodes) {
			<-done
			flight--
			go update(adjustment, &nodes[i])
			flight++
			i++
		}
		for j := 0; j < flight; j++ {
			<-done
		}

		if g.Verbose {
			fmt.Println("computing delta...")
		}
		Δ, leak = 0, 0
		for source := range nodes {
			node := &nodes[source]
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

		if g.Verbose {
			fmt.Println(Δ, ε)
		}
	}

	for key, value := range g.index {
		callback(key, nodes[value].weight[a])
	}
}

// Reset clears all the current graph data.
func (g *Graph32) Reset(size ...int) {
	capacity := 8
	if len(size) == 1 {
		capacity = size[0]
	}
	g.count = 0
	g.index = make(map[uint64]uint, capacity)
	g.nodes = make([]Node32, 0, capacity)
}
