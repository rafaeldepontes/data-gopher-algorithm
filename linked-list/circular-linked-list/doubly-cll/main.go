package main

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyCircularLinkedList struct {
	Head   *Node
	Length int
}

func NewDoublyCLL(data int) *DoublyCircularLinkedList {
	node := &Node{Value: data}
	node.Next = node
	node.Prev = node

	return &DoublyCircularLinkedList{
		Head:   node,
		Length: 1,
	}
}

func (list *DoublyCircularLinkedList) Append(data int) {
	newNode := &Node{Value: data}

	tail := list.Head.Prev

	newNode.Next = list.Head
	newNode.Prev = tail

	tail.Next = newNode
	list.Head.Prev = newNode

	list.Length++
}

func main() {
	list := NewDoublyCLL(1)
	list.Append(2)
	list.Append(3)

	current := list.Head
	iterations := list.Length * 2

	println("========================")
	for range iterations {
		println("Current Node:", current.Value)
		println("Next Node:", current.Next.Value)
		println("Previous Node:", current.Prev.Value)
		println("========================")
		current = current.Next
	}
}
