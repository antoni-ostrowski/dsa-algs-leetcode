package binary_tree

import "fmt"

// search O(n) - linear - data can be anywhere
// no rules for ordering
// nice for representing hierarchies
//        1
//       / \
//      2   3
//     /
//    4

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Value)

	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}
