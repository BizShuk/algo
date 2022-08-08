package linkedlist

// [Notice]:
// [Pattern]: [Get Nth in LinkedList]

func GetNthNode(head *Node, n int) *Node {
	c := &Node{Next: head}
	for c != nil {
		c = c.Next
		n = n - 1
		if n == 0 {
			return c
		}
	}
	return nil
}
