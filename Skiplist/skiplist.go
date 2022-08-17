package skiplist

type ground[K comparable, V any] struct {
	top  *node[K, V]
	item V
}

type node[K comparable, V any] struct {
	prev *node[K, V]
	next *node[K, V]
	down *node[K, V]
	key  K
	grnd *ground[K, V]
}

type Skiplist[K comparable, V any] struct {
	levels int
	head   *node[K, V]
}

func (skl Skiplist[K, V]) findGround(n *node[K, V], key K) *ground[K, V] {
	if n == nil {
		return nil
	}

	result := n
	frontier := n

	for frontier != nil {
		
	}

	if result.down != nil {
		return skl.findGround(result.down, key)
	}

	return result.grnd
}

func (skl Skiplist[K, V]) Find(key K) *V {
	result := skl.findGround(skl.head, key)
	if result == nil {
		return nil
	}

	return &result.item
}

func Make[K comparable, V any](levels int) Skiplist[K, V] {
	//first := new(node[K, V])
	//prev := first
	//for i := 0; i < levels; i++ {
	//	head := new(node[K, V])
	//	prev.next = head
	//	prev = head
	//}

	skl := Skiplist[K, V]{
		levels: levels,
		head:   nil,
	}

	return skl
}

func (skl Skiplist[K, V]) Insert(k K, v V) {

}
