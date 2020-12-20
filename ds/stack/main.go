package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	link  *Node
}

type Stack struct {
	top     int
	topNode *Node
}

func NewStack() *Stack {
	return &Stack{-1, nil}
}

func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Push(n *Node) {
	s.top++
	tmp := new(Node)
	tmp = s.topNode
	s.topNode = n
	n.link = tmp
}

func (s *Stack) Pop() interface{} {
	ret := s.topNode.value
	s.topNode = s.topNode.link
	return ret
}

func (s *Stack) Peek() interface{} {
	return s.topNode.value
}

func main() {
	stack := NewStack()

	stack.Push(&Node{"hi", nil})
	stack.Push(&Node{0, nil})
	stack.Push(&Node{1, nil})

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
