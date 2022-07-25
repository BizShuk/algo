package linkedlist

// Notice: [Pattern], Tip [LinkedList]
// By creating a fake node ahead of 'head' node
// It makes the algo more concise

func Reverse(head *Node) *Node { // Return: tail node
	var prev *Node

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
