package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{
			value: data,
		}
	} else {
		t.insertRecursive(t.root, data)
	}
}

func (t *BinaryTree) Search(data int) *Node {
	return t.searchRecursive(t.root, data)
}

func (t *BinaryTree) searchRecursive(node *Node, data int) *Node {
	if node == nil {
		return nil
	}

	if node.value == data {
		return node
	}

	if node.value < data {
		if n := t.searchRecursive(node.right, data); n != nil {
			return n
		}
	} else {
		if n := t.searchRecursive(node.left, data); n != nil {
			return n
		}
	}

	return nil
}

func (t *BinaryTree) insertRecursive(node *Node, data int) {
	if data > node.value {
		if node.right == nil {
			node.right = &Node{value: data}
			return
		}
		t.insertRecursive(node.right, data)
	} else {
		if node.left == nil {
			node.left = &Node{value: data}
			return
		}
		t.insertRecursive(node.left, data)
	}
}

func main() {
	tree := BinaryTree{}
	tree.Insert(2) // root
	tree.Insert(1) // root -> left
	tree.Insert(5) // root -> right
	tree.Insert(4) // root -> right -> left
	tree.Insert(6) // root -> right -> right

	// println(tree.root.value, tree.root.left.value, tree.root.right.value, tree.root.right.left.value, tree.root.right.right.value)

	if node := tree.Search(6); node != nil {
		println(node.value)
	}

	if node := tree.Search(9); node != nil {
		println(node.value)
	}
}
