package linkedlist

// [Notice]:
// [Pattern]: [Remove Nth in LinkedList]
func RemoveNth(head *Node, n int) *Node {
	p := &Node{Next: head}
	s, f := p, head

	i := 1
	for f != nil {
		if i == n {
			s.Next = f.Next
			return p.Next
		}

		s, f = s.Next, f.Next
		i += 1
	}
	return nil
}

// [Variant]: [Remove Nth in LinkedList] from the end
func RemoveNthFromEnd(head *Node, n int) *Node {
	prev := &Node{Next: head}

	f := prev
	s := prev

	i := 1
	for i <= n {
		if f == nil {
			return nil
		}
		f = f.Next
		i += 1
	}

	for f.Next != nil {
		s = s.Next
		f = f.Next
	}

	s.Next = s.Next.Next
	return prev.Next
}
