package tree

// [Pattern]: [BFS on Tree] Traverse by tree level
func LevelTraverse(root *Node) {
	if root == nil {
		return
	}

	levelStack := []*Node{root}

	for len(levelStack) > 0 {
		nextLevelStack := []*Node{}
		for _, node := range levelStack {
			if node.Left != nil {
				nextLevelStack = append(nextLevelStack, node.Left)
			}
			if node.Right != nil {
				nextLevelStack = append(nextLevelStack, node.Right)
			}
		}
		levelStack = nextLevelStack
	}
}

// [Variant]: [BFS on Tree] Good nodes with iterative
func GoodNodes_iterative(root *Node) int {
	max := -100001
	count := 0

	levelStack := []*Node{root}
	levelMax := []int{max}

	for len(levelStack) > 0 {
		nextLevelStack := []*Node{}
		nextLevelMax := []int{}

		for i, node := range levelStack {
			max = levelMax[i]
			if max <= node.Val {
				max = node.Val
				count += 1
			}

			if node.Left != nil {
				nextLevelStack = append(nextLevelStack, node.Left)
				nextLevelMax = append(nextLevelMax, max)
			}
			if node.Right != nil {
				nextLevelStack = append(nextLevelStack, node.Right)
				nextLevelMax = append(nextLevelMax, max)
			}
		}

		levelStack = nextLevelStack
		levelMax = nextLevelMax
	}

	return count
}

// [Variant]: [BFS on Tree] Good nodes with recursive
func GoodNodes_recursive(root *Node) int {
	return goodNodesFromSub(root, -100001)
}

func goodNodesFromSub(root *Node, max int) int {
	if root == nil {
		return 0
	}

	count := 0

	if max <= root.Val {
		max = root.Val
		count += 1
	}

	count += goodNodesFromSub(root.Left, max)
	count += goodNodesFromSub(root.Right, max)

	return count
}

// [Variant]: [BFS on Tree] Level of Order
func LevelOrder(root *Node) [][]int {

	result := [][]int{}
	if root == nil {
		return result
	}

	levelStack := []*Node{root}

	for len(levelStack) > 0 {
		nextLevel := []*Node{}
		currLevel := []int{}

		for _, node := range levelStack {
			currLevel = append(currLevel, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

		}
		levelStack = nextLevel
		result = append(result, currLevel)
	}

	return result
}

// [Variant]: [BFS on Tree] Right Slice Tree with iterative
func RightSideView_iterative(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	levelStack := []*Node{root}

	for len(levelStack) > 0 {
		var nextLevelStack []*Node
		for _, node := range levelStack {
			if node.Left != nil {
				nextLevelStack = append(nextLevelStack, node.Left)
			}
			if node.Right != nil {
				nextLevelStack = append(nextLevelStack, node.Right)
			}
		}
		result = append(result, levelStack[len(levelStack)-1].Val)
		levelStack = nextLevelStack
	}

	return result
}

// [Variant]: [BFS on Tree] Right Slice Tree with recursive
func RightSideView_recursive(root *Node) []int {
	rightSlice := []int{}
	currLevel := -1

	var traverse func(level int, root *Node)
	traverse = func(level int, root *Node) {
		if root == nil {
			return
		}

		if level > currLevel {
			currLevel = level
			rightSlice = append(rightSlice, root.Val)
		}

		traverse(level+1, root.Right)
		traverse(level+1, root.Left)
	}

	traverse(0, root)

	return rightSlice
}
