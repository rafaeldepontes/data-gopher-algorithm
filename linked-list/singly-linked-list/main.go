package main

type Node struct {
	value int
	next  *Node
}

func NewLinkedList(data int) *Node {
	return &Node{
		value: data,
		next:  nil,
	}
}

func (l Node) NewNode(data int) *Node {
	return &Node{
		value: data,
		next:  nil,
	}
}

func main() {
	sll := NewLinkedList(1)
	sll.next = sll.NewNode(2)

	for sll != nil {
		println("Current Node:", sll.value)
		sll = sll.next
	}
}
