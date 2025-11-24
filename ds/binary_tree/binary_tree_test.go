package binary_tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}

	// Attach a grandchild
	root.Left.Left = &Node{Value: 4}

	// Traverse and print the tree
	fmt.Print("Pre-order Traversal: ")
	PreOrderTraversal(root)
	fmt.Println()
}
