package graph

import (
	"errors"
	"unicode"
)

// https://fodor.org/blog/go-graph/
type Graph struct {
	nodes []*Node
}

type Node struct {
	id      int
	edges   []int
	name    string
	isSmall bool
}

func New() *Graph {
	g := &Graph{
		nodes: []*Node{},
	}

	id := g.AddNode("start")
	g.nodes[id].isSmall = false
	id = g.AddNode("end")
	g.nodes[id].isSmall = false

	return g
}

func (g *Graph) AddNode(name string) (id int) {
	id = len(g.nodes)
	isSmall := IsLower(name)
	g.nodes = append(g.nodes, &Node{
		id:      id,
		edges:   make([]int, 0),
		name:    name,
		isSmall: isSmall,
	})

	return id
}

func (g *Graph) AddEdge(n1, n2 int) {
	g.nodes[n1].edges = append(g.nodes[n1].edges, n2)
	g.nodes[n2].edges = append(g.nodes[n2].edges, n1)
}

func (g *Graph) GetId(name string) (id int, err error) {
	for _, n := range g.nodes {
		if n.name == name {
			return n.id, nil
		}
	}
	return 0, errors.New("node does not exist")
}

func (g *Graph) Neighbours(id int) []int {
	return g.nodes[id].edges
}

func (g *Graph) GetName(id int) string {
	return g.nodes[id].name
}

func (g *Graph) IsSmall(id int) bool {
	return g.nodes[id].isSmall
}

// https://stackoverflow.com/a/59293875
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
