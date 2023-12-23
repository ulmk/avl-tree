package main

import "fmt"

type Node struct {
	key   int
	h     int
	left  *Node
	right *Node
}

func NewNode(key int) *Node {
	return &Node{
		key:   key,
		h:     1,
		left:  nil,
		right: nil,
	}
}

func (t *Node) getH(node *Node) int {
	if node == nil {
		return 0
	}
	return node.h
}

func (t *Node) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	  y
//	x   t3
//
// t1 t2
func (t *Node) RRotate(node *Node) *Node {
	root := node.left
	rootR := root.right
	root.right = node
	node.left = rootR

	node.h = 1 + t.max(t.getH(node.left), t.getH(node.right))
	root.h = 1 + t.max(t.getH(root.left), t.getH(root.right))
	return node
}

//	x
//
// t1   y
//
//	t2 t3
func (t *Node) LRotate(node *Node) *Node {
	root := node.right
	rootL := root.left
	root.left = node
	node.right = rootL

	node.h = 1 + t.max(t.getH(node.left), t.getH(node.right))
	root.h = 1 + t.max(t.getH(root.left), t.getH(root.right))
	return node
}

func (t *Node) getBal(node *Node) int {
	if node == nil {
		return 0
	}
	return t.getH(node.left) - t.getH(node.right)
}

func (t *Node) Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if key < node.key {
		node.left = t.Insert(node.left, key)
	} else if key > node.key {
		node.right = t.Insert(node.right, key)
	} else {
		return node
	}

	node.h = 1 + t.max(t.getH(node.left), t.getH(node.left))

	balance := t.getBal(node)

	if balance > 1 && key < node.left.key {
		return t.RRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return t.LRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = t.LRotate(node.left)
		return t.RRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = t.RRotate(node.right)
		return t.LRotate(node)
	}

	return node
}

func (t *Node) Print(node *Node) {
	if node != nil {
		fmt.Printf("  %v", node.key)
		t.Print(node.left)
		t.Print(node.right)
	}
}
