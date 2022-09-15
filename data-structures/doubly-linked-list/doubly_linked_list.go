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
func (dl *DoublyLinkedList) Find(value uint16) *DoublyLinkedListNode {
	c := dl.Head
	for c != nil {
		if c.Value == value {
			return c
		}
		c = c.Next
	}
	return nil
}
func (dl *DoublyLinkedList) Contains(value uint16) bool {
	return dl.Find(value) != nil
}
func (dl *DoublyLinkedList) Remove(value uint16) bool {
	found := dl.Find(value)
	if found == nil {
		return false
	}
	previous, next := found.Previous, found.Next
	if previous == nil {
		dl.Head = next
		dl.Head.Previous = nil
	} else {
		previous.Next = next
	}
	if next == nil {
		dl.Tail = previous
		dl.Tail.Next = nil
	} else {
		next.Previous = previous
	}
	dl.Count--
	return true

}
func (dl *DoublyLinkedList) ShowAll() {
	c := dl.Head
	for c != nil {
		fmt.Println(c.Value)
		c = c.Next
	}

}

func main() {
	dl := DoublyLinkedList{}
	dl.AddHead(2)
	dl.AddHead(1)
	dl.AddHead(122)
	dl.AddTail(300)
	dl.AddTail(32)
	dl.AddTail(322)

	fmt.Println(dl.Count)
	dl.ShowAll()
	dl.Remove(122)
	dl.Remove(2)
	dl.ShowAll()

}
