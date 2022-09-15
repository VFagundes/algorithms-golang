package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value uint32
}
type Queue struct {
	count uint32
	nodes []*Node
}

func (q *Queue) Enqueue(node *Node) {
	if int(q.count) == len(q.nodes) {
		fmt.Println(errors.New("Queue OverFlow"))
	}
	if int(q.count) == len(q.nodes) {
		q.nodes[len(q.nodes)-1] = node
	} else {
		q.nodes[q.count] = node
	}
	q.count++

}
func (q *Queue) Dequeue() *Node {
	if q.count == 0 {
		fmt.Println(errors.New("Queue Underflow"))
	}
	var n *Node

	if int(q.count) == len(q.nodes) {
		n = q.nodes[0]
		q.nodes[0] = nil
	} else {
		n = q.nodes[q.count-1]
		q.nodes[q.count-1] = nil
	}
	q.count--
	return n
}
func New(capacity int) *Queue {
	n := make([]*Node, capacity)
	return &Queue{nodes: n}
}

func main() {
	q := New(3)
	fmt.Println("queue should be empty")
	fmt.Println(q.count)

	q.Enqueue(&Node{value: 1})
	q.Enqueue(&Node{value: 2})
	q.Enqueue(&Node{value: 3})
	fmt.Println("should print 1, 2, 3")
	for q.count > 0 {
		n := q.Dequeue()
		if n != nil {
			fmt.Println(n.value)
		}
	}
	fmt.Println(len(q.nodes))
	fmt.Println("voila")

}
