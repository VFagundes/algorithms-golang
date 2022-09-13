package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	nodes []*Node
	count uint32
}
type Node struct {
	value uint32
}

func New(capacity int) Stack {
	n := make([]*Node, capacity)
	return Stack{nodes: n}
}

func (s *Stack) Push(item *Node) {
	if int(s.count) == len(s.nodes) {
		fmt.Println(errors.New("error: Stack Overflow"))
	}
	s.nodes[(len(s.nodes)-int(s.count))-1] = item
	s.count++
}
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		fmt.Println(errors.New("error: Stack Underflow"))

	}
	item := s.nodes[len(s.nodes)-int(s.count)]
	s.nodes[len(s.nodes)-int(s.count)] = nil
	s.count--
	return item
}

func main() {
	st := New(3)
	fmt.Println("stack should be empty")
	fmt.Println(st.count)

	st.Push(&Node{value: 3})
	st.Push(&Node{value: 2})
	st.Push(&Node{value: 1})
	fmt.Println("should print 1, 2, 3")
	for st.count > 0 {
		n := st.Pop()
		if n != nil {
			fmt.Println(n.value)
		}
	}
	fmt.Println(len(st.nodes))
	fmt.Println("voila")
}
