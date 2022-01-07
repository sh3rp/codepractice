package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) AddNode(node *Node) {
	if node.Value > n.Value {
		if n.Right != nil {
			n.Right.AddNode(node)
		} else {
			n.Right = node
		}
	} else {
		if n.Left != nil {
			n.Left.AddNode(node)
		} else {
			n.Left = node
		}
	}
}

func Height(node *Node) float64 {
	if node == nil {
		return 0
	}

	return 1 + math.Max(float64(Height(node.Left)), float64(Height(node.Right)))
}

func IsBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	lh := Height(node.Left)
	rh := Height(node.Right)

	if math.Abs(lh-rh) <= 1 && IsBalanced(node.Left) && IsBalanced(node.Right) {
		return true
	}

	return false
}

func main() {
	root := NewNode(4)
	root.AddNode(NewNode(3))
	root.AddNode(NewNode(5))
	root.AddNode(NewNode(2))
	root.AddNode(NewNode(6))
	//root.AddNode(NewNode(7))
	//root.AddNode(NewNode(8))
	//root.AddNode(NewNode(9))

	fmt.Printf("%+v\n", root)
	fmt.Printf("Balanced: %+v\n", IsBalanced(root))
}
