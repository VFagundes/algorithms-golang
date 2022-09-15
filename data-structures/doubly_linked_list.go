package main

import "fmt"

type DoublyLinkedListNode struct {
	Value    uint16
	Next     *DoublyLinkedListNode
	Previous *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head  *DoublyLinkedListNode
	Tail  *DoublyLinkedListNode
	Count uint16
}

func (dl *DoublyLinkedList) AddHead(value uint16) {
	n := &DoublyLinkedListNode{Value: value, Next: dl.Head}
	if dl.Head != nil {
		dl.Head.Previous = n
	}
	dl.Head = n
	if dl.Tail == nil {
		dl.Tail = dl.Head
	}
	dl.Count++

}
func (dl *DoublyLinkedList) AddTail(value uint16) {
	if dl.Tail == nil {
		dl.AddHead(value)
		return
	}
	n := &DoublyLinkedListNode{Value: value, Previous: dl.Tail}
	dl.Tail.Next = n
	dl.Tail = n
	dl.Count++

}

func main() {
	dl := DoublyLinkedList{}
	dl.AddHead(2)
	dl.AddHead(1)
	dl.AddHead(122)
	fmt.Println(dl.Head.Value)
	fmt.Println(dl.Tail.Value)

}
