package main

import "fmt"

func main() {
	fmt.Println("AVL TREE")
	//node := NewNode(10)
	var node *Node
	node = node.Insert(node, 20)
	node = node.Insert(node, 30)
	node = node.Insert(node, 15)
	node = node.Insert(node, 25)
	node = node.Insert(node, 40)
	node = node.Insert(node, 50)
	node.Print(node)
}
