package linkedlist

// Notice: caution for list operartion

type LRUCache struct {
	Cap   int
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	return LRUCache{Cap: capacity, Cache: map[int]*Node{}, Head: head, Tail: head}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.Remove(node)
		this.Put(node.Key, node.Val)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Remove(node *Node) {
	delete(this.Cache, node.Key)

	prev, next := node.Prev, node.Next
	prev.Next = next
	if next != nil {
		next.Prev = prev
	}

	if this.Tail == node {
		this.Tail = prev
	}

}
func (this *LRUCache) Insert(node *Node) {
	this.Tail.Next, node.Prev = node, this.Tail
	this.Tail = node
	this.Cache[node.Key] = node
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{Key: key, Val: value}

	if v, ok := this.Cache[key]; ok { // remove this key if exist
		this.Remove(v)
	} else if len(this.Cache) == this.Cap {
		this.Remove(this.Head.Next)
	}

	this.Insert(node)
}
