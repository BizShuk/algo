package binerytree

// [Notice]: tree node can be nil
// 1. root could nil
func Invert(root *Node) *Node {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	Invert(root.Left)
	Invert(root.Right)

	return root
}
