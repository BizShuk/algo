package tree

// [Notice]: Build tree from preorder and inorder
// PreOrder:
// - 1st elem of preorder is each root
// - 1st part after 1st elem is left tree
// - 2nd part after 1st part is right tree

// InOrder:
// - size of root left is 1st part of PreOrder
// - size of root right is 2nd part of PreOrder

import "math"

func BuildTree(preorder []int, inorder []int) *Node {
	if len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &Node{Val: preorder[0]}
	rootIndex := findIndex(inorder, root.Val)
	// Warning: for boundary
	root.Left = BuildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = BuildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

// inorder slice is sub-slice already
func findIndex(inorder []int, key int) int {
	for i, v := range inorder {
		if v == key {
			return i
		}
	}
	return math.MinInt
}
