package day1

type LRUCache struct {
	capacity, size int
	cache          map[int]*Node
	head, tail     *Node
}
type Node struct {
	key, value int
	prev, next *Node
}

func Init(key, val int) *Node {
	return &Node{key, val, nil, nil}
}
func constructor(capacity int) *LRUCache {
	l := &LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*Node),
		head:     Init(0, 0),
		tail:     Init(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.remove(node)
	l.moveToHead(node)

	return node.value
}
func (l *LRUCache) Put(key, val int) {
	if _, ok := l.cache[key]; !ok {
		node := Init(key, val)
		l.cache[key] = node
		l.moveToHead(node)
		l.size++
		if l.size > l.capacity {

			delete(l.cache, l.tail.prev.key)
			l.remove(l.tail.prev)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = val
		l.remove(node)
		l.moveToHead(node)
	}
}
func (l *LRUCache) moveToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}
func (l *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//func main() {
//	lruCache := constructor(4)
//	lruCache.Put(1, 1)
//	lruCache.Put(2, 2)
//	lruCache.Put(3, 3)
//	lruCache.Put(4, 4)
//	for lruCache.head.next != nil {
//		fmt.Printf("%d ", lruCache.head.key)
//		lruCache.head = lruCache.head.next
//	}
//}
