package bst

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	tree := &BST{}
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	fmt.Print("initial tree [in order]:")
	tree.InOrderTraversal()

	// 2. Search elements
	fmt.Println("Search 40:", tree.Search(40)) // true
	fmt.Println("Search 99:", tree.Search(99)) // false

	// 3. Delete a node with 2 children (30)
	fmt.Println("\nDeleting 30...")
	tree.Delete(30)

	fmt.Print("Tree after deleting 30: ")
	tree.InOrderTraversal()
}
