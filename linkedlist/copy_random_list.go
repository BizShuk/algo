package linkedlist

// Notice: A mapping between old and new is required
// Notice: [Variant] In-Place space
// 1. Chain new node after curr old node
// 2. Find new random node by next node of random node of old node

func CopyRandomList(head *Node) *Node {
	mapping := map[*Node]*Node{}

	preHead := &Node{} // point to new copied head
	var prev, curr, nCurr *Node = preHead, head, nil
	for curr != nil { // create a new copy list with Next pointer only
		prev.Next = &Node{Val: curr.Val}
		mapping[curr] = prev.Next
		prev, curr = prev.Next, curr.Next
	}

	curr, nCurr = head, preHead.Next

	for curr != nil { // mapping the random pointer to the new copied node
		nCurr.Random = mapping[curr.Random]
		curr, nCurr = curr.Next, nCurr.Next
	}

	return preHead.Next
}
