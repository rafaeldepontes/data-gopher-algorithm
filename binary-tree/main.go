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

func (t *BinaryTree) PrintTreePreorder() {
	print("Tree: ")
	t.preorderTraversal(t.root)
	println()
}

func (t *BinaryTree) PrintTreeInorder() {
	print("Tree: ")
	t.inorderTraversal(t.root)
	println()
}

func (t *BinaryTree) PrintTreePostorder() {
	print("Tree: ")
	t.postorderTraversal(t.root)
	println()
}

func (t *BinaryTree) postorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.postorderTraversal(node.left)
	}

	if node.right != nil {
		t.postorderTraversal(node.right)
	}

	print(node.value, " ")
}

func (t *BinaryTree) inorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.inorderTraversal(node.left)
	}

	print(node.value, " ")

	if node.right != nil {
		t.inorderTraversal(node.right)
	}
}

func (t *BinaryTree) preorderTraversal(node *Node) {
	if node == nil {
		return
	}

	print(node.value, " ")
	if node.left != nil {
		t.preorderTraversal(node.left)
	}

	if node.right != nil {
		t.preorderTraversal(node.right)
	}
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

	if node := tree.Search(6); node != nil {
		println("Node:", node.value)
	}

	if node := tree.Search(9); node != nil {
		println("Node:", node.value)
	}

	tree.PrintTreePreorder()
	tree.PrintTreeInorder()
	tree.PrintTreePostorder()

}
