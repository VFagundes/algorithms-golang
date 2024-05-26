package main

import "fmt"

type Node struct {
	Data int32
	Next *Node
}
type LinkedList struct {
	Head   *Node
	Length int32
}

func New(Value int32) *LinkedList {
	return &LinkedList{Head: &Node{Data: Value}, Length: 1}
}

func (l *LinkedList) Prepend(Value int32) int32 {
	next := l.Head
	l.Head = &Node{Data: Value}
	l.Head.Next = next
	l.Length++
	return l.Length
}
func (l *LinkedList) PrintList() {
	current := l.Head
	for {
		fmt.Printf("Node value: %d\n", current.Data)
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	fmt.Printf("LinkedList Length: %d\n", l.Length)
}

func (l *LinkedList) Delete(Value int32) int32 {
	if l.Length == 0 {
		return 0
	}
	if l.Head.Data == Value {
		l.Head = l.Head.Next
		l.Length--
		return l.Length
	}
	previous := l.Head
	//-1 0 1
	for previous.Next != nil && previous.Next.Data != Value {
		previous = previous.Next
	}
	if previous.Next == nil {
		fmt.Println("No values found")
	} else {
		previous.Next = previous.Next.Next
		l.Length--
	}

	return l.Length
}

func (l *LinkedList) InsertAtTheEnd(Value int32) int32 {
	if l.Length == 0 {
		l = New(Value)
		return l.Length
	}

	lastNode := l.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = &Node{Data: Value}
	l.Length++

	return l.Length
}
func (l *LinkedList) Reverse() {
	var prev *Node
	var next *Node
	current := l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev

}

func main() {
	l := New(-1)
	i := 1
	for i = range 4 {
		l.Prepend(int32(i))
	}
	l.PrintList()
	l.Reverse()
	l.PrintList()
	// l.Reverse()

	// l.Prepend(0)
	// l.Prepend(-1)
	// l.PrintList()
	// l.Delete(1)
	// l.Delete(12)
	// l.Delete(111)
	// l.InsertAtTheEnd(2)

	fmt.Println("print linked list code")
}
