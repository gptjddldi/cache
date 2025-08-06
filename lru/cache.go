package lru

type Node struct {
	Next *Node
	Prev *Node
	Val  int
	Key  int
}

type LRUCache struct {
	head     *Node
	tail     *Node
	capacity int
	mp       map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{Key: 0, Val: 0}, &Node{Key: 0, Val: 0}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		mp:       make(map[int]*Node),
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.mp[key]; ok {
		c.remove(node)
		c.insert(node)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.mp[key]; ok {
		c.remove(c.mp[key])
	}
	if len(c.mp) == c.capacity {
		c.remove(c.tail.Prev)
	}
	newNode := &Node{Key: key, Val: value}
	c.insert(newNode)
}

func (c *LRUCache) remove(node *Node) {
	delete(c.mp, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) insert(node *Node) {
	c.mp[node.Key] = node
	next := c.head.Next
	next.Prev = node

	node.Next = next
	node.Prev = c.head
	c.head.Next = node
}
