package binerytree

// Notice
// 1. root could nil

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
