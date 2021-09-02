package hash

type node struct {
	key   string
	value interface{}
	next  *node
	//flag  int64
}

func newNodeHead() *node {
	return &node{
		key:   "",
		value: "",
		next:  nil,
	}
}

func (n *node) create(k string, v interface{}) *node {
	if n == nil {
		n = newNodeHead()
	}
	n.key = k
	n.value = v
	return n
}

func (n *node) get(k string) (interface{}, bool) {
	if n.next == nil {
		return nil, false
	}
	for n.next != nil {
		if n.next.key == k {
			return n.next.value, true
		} else {
			n = n.next
		}
	}
	return nil, false
}

func (n *node) add(k string, v interface{}) {
	if _, ok := n.get(k); ok {
		return
	}
	for n.next != nil {
		n = n.next
	}
	n.next = n.next.create(k, v)
}

func (n *node) erase(k string) (interface{}, bool) {
	if n.next == nil {
		return nil, false
	}
	for n.next != nil {
		if n.next.key == k {
			n.next = n.next.next
			return n.next.value, true
		} else {
			n = n.next
		}
	}
	return nil, false
}

type LightyHashMap struct {
	buckets []*node
	//bitLock byte
}

func (h *LightyHashMap) Insert(k string, v interface{}) {
	n := hashNum(k)
	h.buckets[n].add(k, v)
}

func (h *LightyHashMap) Get(k string) (interface{}, bool) {
	n := hashNum(k)
	return h.buckets[n].get(k)
}

func (h *LightyHashMap) Erase(k string) (interface{}, bool) {
	n := hashNum(k)
	return h.buckets[n].erase(k)
}

func hashNum(k string) int {
	index := 0
	for i := 0; i < len(k); i++ {
		index *= (1103515245 + int(k[i]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}
