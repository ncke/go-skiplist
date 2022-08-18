package skiplist

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

// Represents the stored key-value pair.
type ground[K constraints.Ordered, V any] struct {
	key K
	val V
}

// Represents a skip list node that points to the next node at
// the same level, and down to the equivalent node at the
// level below.
type node[K constraints.Ordered, V any] struct {
	next *node[K, V]
	down *node[K, V]
	grnd *ground[K, V]
}

// Skiplist represents a skip list with at most a given number of levels.
type Skiplist[K constraints.Ordered, V any] struct {
	levels int
	head   *node[K, V]
}

// Returns a slice of nodes at descending levels representing each
// of which has a key that is ordered before the target key.
func (skl Skiplist[K, V]) findSteps(n *node[K, V], key K) []*node[K, V] {
	result := n
	frontier := n.next

	for frontier != nil {
		if key < frontier.grnd.key {
			// The frontier has gone beyond the target key, so we have
			// a result at this level.
			break
		}

		result = result.next
		frontier = frontier.next
	}

	arr := []*node[K, V]{result}

	if result.down == nil {
		// This is the base case, there are no lower nodes.
		return arr
	}

	// Re-find at the lower level.
	return append(arr, skl.findSteps(result.down, key)...)
}

// Make constructs a skip list.
func Make[K constraints.Ordered, V any](levels int) Skiplist[K, V] {
	var p *node[K, V] = nil
	for i := 0; i < levels; i++ {
		n := new(node[K, V])
		n.down = p
		p = n
	}

	skl := Skiplist[K, V]{
		levels: levels,
		head:   p,
	}

	return skl
}

// Find yields the value held by the skip list for the given key,
// returns nil if not found.
func (skl Skiplist[K, V]) Find(key K) *V {
	steps := skl.findSteps(skl.head, key)
	last := steps[len(steps)-1]

	if last.grnd == nil || last.grnd.key != key {
		return nil
	}

	return &last.grnd.val
}

// Randomly determine a height for a new node column.
func (skl Skiplist[K, V]) randomHeight() int {
	h := 0
	for {
		h++
		if rand.Float64() < 0.5 || h == skl.levels {
			break
		}
	}

	return h
}

// Insert places a value into the skip list with the given key.
func (skl Skiplist[K, V]) Insert(key K, val V) {
	grnd := new(ground[K, V])
	grnd.key = key
	grnd.val = val

	steps := skl.findSteps(skl.head, key)
	j := len(steps)

	if steps[j-1].grnd != nil && steps[j-1].grnd.key == key {
		steps[j-1].grnd.val = val
		return
	}

	var p *node[K, V] = nil
	for i := 0; i < skl.randomHeight(); i++ {
		j--

		n := new(node[K, V])

		if i == 0 {
			n.grnd = grnd
		}

		n.down = p
		p = n
		n.next = steps[j].next
		steps[j].next = n
	}
}
