package livetest

type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

type LRUCache struct {
	cache    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	capacity int
}

func Remove(n *LRUNode) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (this *LRUCache) AddFirst(n *LRUNode) {
	temp := this.head.next
	this.head.next = n
	n.prev = this.head
	n.next = temp
	temp.prev = n
}

func (this *LRUCache) RemoveLast() {
	temp := this.tail.prev
	this.tail.prev = temp.prev
	temp.prev.next = this.tail
}

func LRUConstructor(capacity int) LRUCache {
	m := LRUCache{}
	m.capacity = capacity
	m.cache = make(map[int]*LRUNode)
	m.head = &LRUNode{}
	m.tail = &LRUNode{}
	m.head.next = m.tail
	m.tail.prev = m.head
	return m
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if ok {
		Remove(v)
		this.AddFirst(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	temp := &LRUNode{
		key:   key,
		value: value,
	}
	v, ok := this.cache[key]
	if ok {
		Remove(v)
		delete(this.cache, key)
	}
	this.AddFirst(temp)
	this.cache[key] = temp

	if len(this.cache) > this.capacity {
		delete(this.cache, this.tail.prev.key)
		this.RemoveLast()

	}
}
