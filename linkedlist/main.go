package main

import "fmt"

type Node struct {
	value int
	link  *Node
}

func (n *Node) Append(node *Node) *Node {
	n.link = node
}

func main() {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}

	n3.append(n4)
	n2.append(n3)
	n1.append(n2)

	fmt.Println(n1.value)                // 1
	fmt.Println(n1.link.value)           // 2
	fmt.Println(n1.link.link.value)      // 3
	fmt.Println(n1.link.link.link.value) // 4
}
