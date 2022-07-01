package bst

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Root() *Node {
	return bst.root
}

func (bst *BST) InsertNode(d int) bool {
	newNode := &Node{data: d}

	if bst.root == nil {
		bst.root = newNode
		return true
	}

	temp := bst.root
	for temp != nil {
		if temp.data > d {
			if temp.left != nil {
				temp = temp.left
			} else {
				temp.left = newNode
				break
			}
		} else if temp.data < d {
			if temp.right != nil {
				temp = temp.right
			} else {
				temp.right = newNode
				break
			}
		} else {
			return false
		}
	}
	return true
}

func (bst *BST) InOrder(root *Node) {
	if root == nil {
		return
	}
	bst.InOrder(root.left)
	fmt.Println(root.data)
	bst.InOrder(root.right)
}

func (bst *BST) PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	bst.InOrder(root.left)
	bst.InOrder(root.right)
}

func (bst *BST) PostOrder(root *Node) {
	if root == nil {
		return
	}
	bst.InOrder(root.left)
	bst.InOrder(root.right)
	fmt.Println(root.data)
}