package fifo

type Cache struct {
	capacity int
	queue    []int
	mp       map[int]int
}

func NewCache(capacity int) Cache {
	return Cache{
		capacity: capacity,
		queue:    []int{},
		mp:       map[int]int{},
	}
}

func (c *Cache) Get(key int) int {
	if _, ok := c.mp[key]; ok {
		return c.mp[key]
	}
	return -1
}

func (c *Cache) Put(key, value int) {
	if _, ok := c.mp[key]; ok {
		c.mp[key] = value
		return
	}
	if len(c.queue) == c.capacity {
		delete(c.mp, c.queue[0])
		c.queue = c.queue[1:]
	}
	c.mp[key] = value
	c.queue = append(c.queue, key)
}
