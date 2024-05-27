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
	// nil<-1<-2<-3
	current := l.Head
	var prev, next *Node
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next

	}
	l.Head = prev

}

func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	// https://leetcode.com/problems/add-two-numbers/
	tmp := &Node{}
	current := tmp
	var carry int32 = 0
	currentl1 := l1.Head
	currentl2 := l2.Head
	var length int32 = 0
	for currentl1 != nil || currentl2 != nil {
		var sum int32 = carry
		if currentl1 != nil {
			sum += currentl1.Data
			currentl1 = currentl1.Next
		}
		if currentl2 != nil {
			sum += currentl2.Data
			currentl2 = currentl2.Next
		}
		current.Next = &Node{Data: sum % 10}
		current = current.Next
		carry = sum / 10
		length++
	}
	if carry > 0 {
		current.Next = &Node{Data: carry}
	}
	l := LinkedList{}
	l.Head = tmp.Next
	l.Length = length
	return &l

}
func runAddTwoNumbers() {
	l1 := New(3)
	l1.Prepend(4)
	l1.Prepend(2)
	l1.PrintList()
	l2 := New(4)
	l2.Prepend(6)
	l2.Prepend(5)
	l2.PrintList()
	addTwoNumbers(l1, l2).PrintList()

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
	fmt.Println("------- end of method execution -------")
	runAddTwoNumbers()

	fmt.Println("print linked list code")
}
