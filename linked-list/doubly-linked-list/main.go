package main

type DoublyLinkedList struct {
	Value int
	Next  *DoublyLinkedList
	Prev  *DoublyLinkedList
}

func NewLinkedList(data int) *DoublyLinkedList {
	return &DoublyLinkedList{
		Value: data,
		Next:  nil,
		Prev:  nil,
	}
}

func (list *DoublyLinkedList) NewNode(data int, prev *DoublyLinkedList) *DoublyLinkedList {
	return &DoublyLinkedList{
		Value: data,
		Next:  nil,
		Prev:  prev,
	}
}

func main() {
	dll := NewLinkedList(1)
	dll.Next = dll.NewNode(2, dll)

	println("======================")
	for dll != nil {
		println("Current Node:", dll.Value)
		println("It has a previous node?", dll.Prev != nil)
		if dll.Prev != nil {
			println("Previous Node:", dll.Prev.Value)
		}
		println("======================")
		dll = dll.Next
	}
}
