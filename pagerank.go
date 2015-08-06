/*
Package pagerank implements the *weighted* PageRank algorithm.
*/
package pagerank

import (
	"math"
)

type _Node struct {
	Weight   float64
	Outbound float64
}

type Graph struct {
	Edges map[int](map[int]float64) // @TODO: This data structure is not ideal.
	Nodes map[int]*_Node
}

// New initializes and returns a new graph.
func New() *Graph {
	return &Graph{
		Edges: make(map[int](map[int]float64)),
		Nodes: make(map[int]*_Node),
	}
}

// Link creates a weighted edge between a source-target node pair.
// If the edge already exists, the weight is incremented.
func (self *Graph) Link(source, target int, weight float64) {
	if _, ok := self.Nodes[source]; ok == false {
		self.Nodes[source] = &_Node{
			Weight:   0,
			Outbound: 0,
		}
	}

	self.Nodes[source].Outbound += weight

	if _, ok := self.Nodes[target]; ok == false {
		self.Nodes[target] = &_Node{
			Weight:   0,
			Outbound: 0,
		}
	}

	if _, ok := self.Edges[source]; ok == false {
		self.Edges[source] = map[int]float64{}
	}

	self.Edges[source][target] += weight
}

// Rank computes the PageRank of every node in the directed graph.
// α (alpha) is the damping factor, usually set to 0.85.
// ε (epsilon) is the convergence criteria, usually set to a tiny value.
//
// This method will run as many iterations as needed, until the graph converges.
func (self *Graph) Rank(α, ε float64, callback func(id int, rank float64)) {
	Δ := float64(1.0)
	inverse := 1 / float64(len(self.Nodes))

	// Normalize all the edge weights so that their sum amounts to 1.
	for source := range self.Edges {
		if self.Nodes[source].Outbound > 0 {
			for target, _ := range self.Edges[source] {
				self.Edges[source][target] /= self.Nodes[source].Outbound
			}
		}
	}

	for key := range self.Nodes {
		self.Nodes[key].Weight = inverse
	}

	for Δ > ε {
		leak := float64(0)
		nodes := map[int]float64{}

		for key, value := range self.Nodes {
			nodes[key] = value.Weight

			if value.Outbound == 0 {
				leak += value.Weight
			}

			self.Nodes[key].Weight = 0
		}

		leak *= α

		for source := range self.Nodes {
			for target := range self.Edges[source] {
				self.Nodes[target].Weight += α * nodes[source] * self.Edges[source][target]
			}

			self.Nodes[source].Weight += (1 - α) * inverse + leak * inverse
		}

		Δ = 0

		for key, value := range self.Nodes {
			Δ += math.Abs(value.Weight - nodes[key])
		}
	}

	for key, value := range self.Nodes {
		callback(key, value.Weight)
	}
}

// Reset clears all the current graph data.
func (self *Graph) Reset() {
	self.Edges = make(map[int](map[int]float64))
	self.Nodes = make(map[int]*_Node)
}
