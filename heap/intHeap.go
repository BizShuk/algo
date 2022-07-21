package heap

import (
	"sort"
)

// TODO: How to integrate sort.IntSlice for IntHeap
type IntHeap struct {
	sort.Interface
}

func (h *IntHeap) Push(x any) {
	w := sortSlice(*h)
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
// func main() {
// 	h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }
