package bst

import "fmt"

// Left Subtree Rule: The values of all nodes in the left subtree of a node must be strictly less than the node’s value.
// Right Subtree Rule: The values of all nodes in the right subtree of a node must be strictly greater than the node’s value.
// No Duplicates: Standard BSTs do not allow duplicate keys (though variations can store duplicates by appending them to the left or right consistently).
// In-Order Sorting: Running an In-Order traversal (Left → Root → Right) on a BST will always output the elements in perfectly sorted ascending order.

// time complexity
// all del/insert/ are O(logn) on avarage (balanced), worst is O(n)
// traversals depend searhc is the same as ops, inOrder is always O(n) cause it goes through everything

// use cases:
// - dynamic sorted lists, inserting item inserts it to correct (sorted) place
// - good for range queries, data is structurarly sorted, BST can easly find all elements withing specific range (like between >20 & <50)
// - priority - we can quickly find closest higher/lowest value to given target, so like priority based queue

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) insert(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.insert(val)
		}
	} else if val > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.insert(val)
		}
	}
}
func (n *Node) search(target int) bool {
	if n == nil {
		return false
	}
	if target == n.Value {
		return true
	}
	if target < n.Value {
		return n.Left.search(target)
	}
	return n.Right.search(target)
}

func (n *Node) delete(val int) *Node {
	if n == nil {
		return nil
	}
	if val < n.Value {
		n.Left = n.Left.delete(val)
	} else if val > n.Value {
		n.Right = n.Right.delete(val)
	} else {
		// found node

		// case: 1 children
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// case: 2 children
		// Why findMin of right? That node is guaranteed: bigger than everything in left subtree, smaller than everything else in right subtree. So copying it
		// keeps the Left/Right rule intact. And by definition it has no left child, so deleting it is easy case 1 or 2.
		minNode := n.Right.findMin()
		n.Value = minNode.Value
		n.Right = n.Right.delete(minNode.Value)

	}
	return n

}
func (n *Node) findMin() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}
func (n *Node) inOrder() {
	if n == nil {
		return
	}
	n.Left.inOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.inOrder()
}

type BST struct {
	Root *Node
}

func (b *BST) Insert(val int) {
	if b.Root == nil {
		b.Root = &Node{Value: val}
		return
	}
	b.Root.insert(val)
}

func (b *BST) Search(val int) bool {
	return b.Root.search(val)
}

func (b *BST) Delete(val int) {
	b.Root = b.Root.delete(val)
}
func (b *BST) InOrderTraversal() {
	b.Root.inOrder()
	fmt.Println()

}
