package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	// pq[i].index = i
	// pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	// n := len(*pq)
	item := x.(*Item)
	// item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	// heap.Fix(pq, item.index)
}

func topKFrequent(nums []int, k int) []int {
	freqCounter := make(map[int]int)

	for _, val := range nums {
		freqCounter[val] += 1
	}

	// need a Heap, priority queue, or monotonic stack
	// if get top k, use min heap and pop if heap over k
	pq := make(PriorityQueue, len(freqCounter))

	i := 0
	for k, v := range freqCounter {
		pq[i] = &Item{value: k, priority: v, index: i}
		i += 1
	}

	heap.Init(&pq)

	for len(pq) != k {
		popd := heap.Pop(&pq)
		fmt.Println(popd)
	}

	result := make([]int, len(pq))

	for i := 0; len(pq) != 0; i++ {
		result[i] = heap.Pop(&pq).(*Item).value
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
