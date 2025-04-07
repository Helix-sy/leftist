package leftist

import (
	"gitlab.lrz.de/hm/goal-core/base"
	"golang.org/x/exp/constraints"
)

// A LeftistHeap is an implementation of collection.PriorityQueue
type LeftistHeap[E any, K constraints.Ordered] struct {
	min  *leftistNode[E, K]
	size uint
}

type leftistNode[E any, K constraints.Ordered] struct {
	element         E
	key             K
	dist            uint
	left, right, up *leftistNode[E, K]
}

func NewLeftistHeap[E any, K constraints.Ordered]() *LeftistHeap[E, K] {
	return &LeftistHeap[E, K]{min: nil}
}

func (heap *LeftistHeap[E, V]) IsEmpty() bool {
	return heap.min == nil
}

func (heap *LeftistHeap[E, K]) Size() uint {
	return heap.size
}

// Internal operation to merge two leftist trees given as root nodes.
// Note: Both arguments may be nil (in which case the result will be nil, too)
// This method will not compute any sizes, but correct the dist-Values along the merging
// path and flip left/right where needed
func (node *leftistNode[E, K]) merge(other *leftistNode[E, K]) *leftistNode[E, K] {
	panic("implement me")
}

// Insert adds a new element with a given priority (key).
func (heap *LeftistHeap[E, K]) Insert(element E, key K) base.Handle {
	// newNode := &leftistNode[E, K]{element: element, key: key, dist: 1}
	// ...
	panic("implement me")
}

// ExtractMin extracts the minimum element (with key/priority).
func (heap *LeftistHeap[E, K]) ExtractMin() (E, K) {
	if heap.IsEmpty() {
		panic("leftist heap empty")
	}
	panic("implement me")
}

// Remove removes an element (by handle) from the heap,
// splitting it in three parts, fixing one of them and
// merging them to one tree again.
func (heap *LeftistHeap[E, K]) Remove(handle base.Handle) {
	//node, ok := handle.(*leftistNode[E, K])
	//if !ok {
	//	panic("bad handle type")
	// }
	//panic("implement me")
}

// DecreaseKey reduces the key of an element.
func (heap *LeftistHeap[E, K]) DecreaseKey(handle base.Handle, key K) {
	node, ok := handle.(*leftistNode[E, K])
	if !ok {
		panic("bad handle type")
	}
	if node.key <= key {
		panic("called DecreaseKey() with higher key.")
	}
	panic("implement me")
}
