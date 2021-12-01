package main

import "fmt"

type Node struct {
	Id          int
	Adjacencies []*Node
}

func (n *Node) IsCyclic() bool {
	visited := make(map[int]bool)

	return DFS(n, visited)
}

func DFS(n *Node, visited map[int]bool) bool {
	if n == nil {
		return false
	}

	if visited[n.Id] {
		return true
	}

	visited[n.Id] = true

	for _, adj := range n.Adjacencies {
		return DFS(adj, visited)
	}

	return false
}

func (n *Node) Connect(node *Node) {
	node.Adjacencies = append(node.Adjacencies, node)
}

func NewNode(value int) *Node {
	return &Node{value, []*Node{}}
}

func main() {

	// Cycle: false

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	n1.Connect(n2)
	n2.Connect(n3)

	fmt.Printf("Cycle: %+v\n", n1.IsCyclic())

	// Cycle: true

	m1 := NewNode(1)
	m2 := NewNode(2)
	m3 := NewNode(3)

	m1.Connect(m2)
	m2.Connect(m3)
	m3.Connect(m1)

	fmt.Printf("Cycle: %+v\n", m1.IsCyclic())
}
