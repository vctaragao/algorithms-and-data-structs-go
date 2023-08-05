package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

type queue []*Node

func (q *queue) pop() *Node {
	nodeQueue := *q

	first := nodeQueue[0]
	nodeQueue = nodeQueue[1:]

	*q = nodeQueue

	return first
}

func (n *Node) inOrderTraverse() {
	if n == nil {
		return
	}

	n.left.inOrderTraverse()
	fmt.Println(n.data)
	n.right.inOrderTraverse()
}

func (n *Node) preOrderTraverse() {
	if n == nil {
		return
	}

	fmt.Println(n.data)
	n.left.preOrderTraverse()
	n.right.preOrderTraverse()
}

func (n *Node) postOrderTraverse() {
	if n == nil {
		return
	}

	n.left.postOrderTraverse()
	n.right.postOrderTraverse()

	fmt.Println(n.data)
}

func (n *Node) levelOrderTraverse() []string {

	nodes := []string{}
	nodeQueue := queue{n}

	for len(nodeQueue) > 0 {
		node := nodeQueue.pop()
		nodes = append(nodes, node.data)

		if node.left != nil {
			nodeQueue = append(nodeQueue, node.left)

			if node.right != nil {
				nodeQueue = append(nodeQueue, node.right)
			}
		}
	}

	return nodes
}

func main() {

	n3 := &Node{
		data: "first grandchild node",
	}

	n4 := &Node{
		data: "second grandchild node",
	}

	n1 := &Node{
		data:  "first child node",
		left:  n3,
		right: n4,
	}

	n2 := &Node{
		data: "second child node",
	}

	n0 := &Node{
		data:  "root node",
		left:  n1,
		right: n2,
	}

	fmt.Println("In-Order Traverse")
	n0.inOrderTraverse()

	fmt.Println("")

	fmt.Println("Pre-Order Traverse")
	n0.preOrderTraverse()

	fmt.Println("")

	fmt.Println("Post-Order Traverse")
	n0.postOrderTraverse()

	fmt.Println("")

	fmt.Println("Level-Order Traverse")
	fmt.Println(n0.levelOrderTraverse())
}
