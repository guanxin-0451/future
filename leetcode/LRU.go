package leetcode

type LRUCache struct {
	head       *Node
	tail       *Node
	dict       map[int]*Node
	usedVolume int
	maxVolume  int
}

type Node struct {
	key   int
	value int
	next  *Node
	pre   *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		head:      CreateNode(0, 0),
		tail:      CreateNode(0, 0),
		dict:      make(map[int]*Node),
		maxVolume: capacity,
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.pre = lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	_, ok := this.dict[key]
	if !ok {
		return -1
	}

	this.reRank(this.dict[key])
	return this.dict[key].value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.dict[key]
	if ok {
		this.dict[key].value = value
		this.reRank(this.dict[key])
	} else {
		node := CreateNode(key, value)
		this.add(node)
		this.usedVolume++
		this.dict[key] = node
		if this.usedVolume > this.maxVolume {
			tail := this.removeTail()
			delete(this.dict, tail.key)
			this.usedVolume--
		}
	}

}

func CreateNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (c *LRUCache) removeTail() *Node {
	tail := c.tail.pre
	c.remove(c.tail.pre)

	return tail
}

func (c *LRUCache) add(node *Node) {
	node.next = c.head.next
	node.pre = c.head
	c.head.next = node
	node.next.pre = node
}

func (c *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (c *LRUCache) reRank(node *Node) { // 重排
	c.remove(node)
	c.add(node)
}
