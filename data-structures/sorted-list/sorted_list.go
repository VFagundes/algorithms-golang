package main

type SortedList struct {
	Count int
	node  *Node
	Keys []*Node
}
type Node struct {
	Value int
	next  *Node
}

func (sl *SortedList) Add(value int) {
	if sl.node == nil {
		sl.node = &Node{Value: value}
		return
	}else if 

}
func (sl *SortedList) Remove(value int) bool {
	return false
}
func (sl *SortedList) GetKey(value int) *Node {
	return nil
}

func main() {

}
