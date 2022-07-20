package linkedlist

func Reverse(head *Node) *Node {
	var prev *Node

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
