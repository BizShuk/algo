package tree

func IsValidBST(root *Node) bool {
	if root == nil {
		return true
	}

	lastValAssigned := false

	lastVal := 0
	curr := root
	for curr != nil {
		lastVal = curr.Val
		curr = curr.Left
	}

	var dfs func(root *Node) bool

	dfs = func(root *Node) bool {
		if root == nil {
			return true
		}

		valid := dfs(root.Left)

		if !lastValAssigned {
			lastVal = root.Val
			lastValAssigned = true
		} else if lastVal >= root.Val || !valid {
			return false
		}

		lastVal = root.Val

		return dfs(root.Right)
	}
	return dfs(root)
}
