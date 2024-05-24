package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	Length int
	Head   *Node
}

func (l *LinkedList) Prepend(n *Node) {
	sec := l.Head
	l.Head = n
	n.next = sec
	l.Length++
}
func (l *LinkedList) Delete(value int) {
	if l.Length == 0 {
		fmt.Println("empty list")
		return
	}
	previous := l.Head
	if previous.data == value {
		l.Head = previous.next
		l.Length--
		return
	}
	for previous.next.data != value {
		previous = previous.next
		if previous.next == nil {
			fmt.Println("value not found")
			return
		}
	}
	previous.next = previous.next.next
	//1,2,3,4 param:3
	//
	l.Length--
}

func New(data int) LinkedList {
	return LinkedList{Head: &Node{data: data}, Length: 1}
}
func (l *LinkedList) Print() {
	current := l.Head
	for {
		fmt.Printf("value: %d list size: %d\n", current.data, l.Length)

		current = current.next
		if current == nil {
			break
		}

	}
}
func main() {
	l := New(1)
	l.Prepend(&Node{data: 2})
	l.Prepend(&Node{data: 3})
	l.Prepend(&Node{data: 4})
	l.Prepend(&Node{data: 5})
	l.Prepend(&Node{data: 6})

	l.Print()
	fmt.Println("------")
	// l.Delete(3)
	l.Delete(6)
	// l.Delete(4)
	l.Print()
}
