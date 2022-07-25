package linkedlist

func reorderList(head *Node) {
	fakeHead := &Node{Next: head}

	f, s := fakeHead, fakeHead
	for f != nil && f.Next != nil { // Should be a cycle linked list
		s = s.Next
		f = f.Next.Next
	}

	// reverse from s
	h, t := head, Reverse(s.Next)
	s.Next = nil

	merged := fakeHead
	for h != nil && t != nil { // how h t know they are bypass the other => set its Next to nil
		merged.Next = h
		merged = merged.Next
		h = h.Next // Notice: [LinkedList] might mistake by change node next in sequence operations

		merged.Next = t
		merged = merged.Next
		t = t.Next
	}

	if h != nil {
		merged.Next = h
	}
}
