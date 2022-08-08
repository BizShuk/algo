package linkedlist

// [Notice]: [Variant] findDuplicateNum with array

func HasCycle(head *Node) bool {
	s, f := head, head
	for f != nil && f.Next != nil { // [Notice]: [Fast-Slow pattern] f!= nil && f.Next != nil
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}
