package bst

import (
	"fmt"
	"testing"
)

func TestInsertNode(t *testing.T) {

	bst := BST{}
	tests := []int{5, 8, 3, 12, 9}
	for _,v := range tests {
		if ok:=bst.InsertNode(v); !ok {
			t.Error("FAILED.")
		} else {
			t.Log("PASSED.")
		}
	}
	
	fmt.Println("InOrder Traversal")
	bst.InOrder(bst.root)
	fmt.Println("PreOrder Traversal")
	bst.PreOrder(bst.root)
	fmt.Println("PostOrder Traversal")
	bst.PostOrder(bst.root)
}
