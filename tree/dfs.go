package tree

// [Pattern]: [DFS on tree] withrecursive and iterative
// Could go left/right
func Dfs_recursive(root *Node) {
	if root == nil {
		return
	}
	// Do something in pre-order
	Dfs_recursive(root.Left)
	// Do something in in-order
	Dfs_recursive(root.Right)
	// Do something in post-order
}

// No post-order possible
func Dfs_iterative(root *Node) {
	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			// Do something in pre-order
			stack = append(stack, curr)
			curr = curr.Left
		}

		popd := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Do something in in-order
		curr = popd.Right
	}
}

// [Variant]: [DFS on tree] kth smaleest num
func KthSmallest_iterative(root *Node, k int) int {
	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		popd := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k -= 1
		if k == 0 {
			return popd.Val
		}

		curr = popd.Right
	}
	return -1
}

func KthSmallest_recursive(root *Node, k int) int {
	kthVal := 0

	var dfs func(root *Node)

	dfs = func(root *Node) {
		if root == nil {
			return
		}

		dfs(root.Left)
		k -= 1
		if k == 0 {
			kthVal = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return kthVal
}
