package linkedlist

// Notice: [Important] Caution of LinkedList op. Just like fakeHead, fakeTail works in doulbe LinkedList

type LRUCache struct {
	Cap   int
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Prev = tail, head
	return LRUCache{Cap: capacity, Cache: map[int]*Node{}, Head: head, Tail: tail}
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
	key := node.Key
	if _, ok := this.Cache[key]; !ok {
		return
	}

	node = this.Cache[key]
	delete(this.Cache, key)

	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

func (this *LRUCache) Insert(node *Node) {
	prev := this.Tail.Prev

	node.Next, this.Tail.Prev = this.Tail, node
	prev.Next, node.Prev = node, prev

	this.Cache[node.Key] = node
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{Key: key, Val: value}

	this.Remove(node)
	this.Insert(node)

	if len(this.Cache) > this.Cap {
		this.Remove(this.Head.Next)
	}
}
