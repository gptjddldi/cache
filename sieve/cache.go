package sieve

type Sieve struct {
	head     *Node
	tail     *Node
	mp       map[int]*Node
	hand     *Node
	capacity int
}

type Node struct {
	Next *Node
	Prev *Node

	Visited bool
	Val     int
	Key     int
}

func NewSieve(capacity int) Sieve {
	head, tail := &Node{Key: 0, Val: 0}, &Node{Key: 0, Val: 0}

	head.Next = tail
	tail.Prev = head

	return Sieve{
		head:     head,
		tail:     tail,
		mp:       map[int]*Node{},
		capacity: capacity,
	}
}

func (c *Sieve) Get(key int) int {
	if v, ok := c.mp[key]; ok {
		v.Visited = true
		return v.Val
	}
	return -1
}

func (c *Sieve) Put(key, value int) {
	if node, ok := c.mp[key]; ok {
		node.Val = value
		return
	}

	if len(c.mp) >= c.capacity {
		c.evict()
	}

	newNode := &Node{
		Key: key,
		Val: value,
	}
	c.add(newNode)
}

func (c *Sieve) evict() {
	if c.hand == nil {
		c.hand = c.tail.Prev
	}

	for c.hand != c.head && c.hand.Visited {
		c.hand.Visited = false
		c.hand = c.hand.Prev
	}

	if c.hand == c.head {
		c.hand = c.tail.Prev
		for c.hand != c.head && c.hand.Visited {
			c.hand.Visited = false
			c.hand = c.hand.Prev
		}
	}

	evictNode := c.hand
	c.hand = evictNode.Prev
	c.remove(evictNode)
}

// 새로운 노드를 리스트의 맨 앞에 추가
func (c *Sieve) add(node *Node) {
	node.Next = c.head.Next
	node.Prev = c.head
	c.head.Next.Prev = node
	c.head.Next = node
	c.mp[node.Key] = node
}

// 노드를 리스트에서 제거
func (c *Sieve) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(c.mp, node.Key)
}
