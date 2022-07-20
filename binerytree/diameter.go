package binerytree

// Notice
// 1. two vaule need to take care, diameter in current tree and longest path from this node
//    Use inner function with outside variable to deal the issue
//    or, return two values in each function call
// 2. the count is node count, not diameter
// 3. hard to use iterative way to implement

func diameterOfBinaryTree(root *Node) int {
	maxDiameter := 0

	var dfs func(root *Node) int

	dfs = func(root *Node) int /* current lognest in sub-trees */ {
		if root == nil {
			return 0
		}

		leftDiameter := dfs(root.Left)
		rightDiameter := dfs(root.Right)
		if leftDiameter+rightDiameter+1 > maxDiameter {
			maxDiameter = leftDiameter + rightDiameter + 1
		}

		return Max(leftDiameter, rightDiameter) + 1
	}

	dfs(root)
	return maxDiameter - 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
