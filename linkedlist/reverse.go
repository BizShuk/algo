package linkedlist

// [Pattern]: [Reverse LinkedList]
func Reverse(head *Node) *Node { // Return: tail node
	var prev *Node // [Notice]: A fake node will help on code tidiness

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
