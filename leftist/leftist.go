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
	if node == nil {
		return other
	}
	if other == nil {
		return node
	}

	// Ensure min-heap property
	if other.key < node.key {
		node, other = other, node
	}

	// Recursively merge right subtree with other tree
	node.right = node.right.merge(other)

	// Update parent pointer
	if node.right != nil {
		node.right.up = node
	}

	// Maintain leftist property
	if node.left == nil {
		node.left, node.right = node.right, node.left
	} else if node.right != nil && node.left.dist < node.right.dist {
		node.left, node.right = node.right, node.left
	}

	// Update dist value
	if node.right == nil {
		node.dist = 1
	} else {
		node.dist = node.right.dist + 1
	}

	return node
}

// Insert adds a new element with a given priority (key).
func (heap *LeftistHeap[E, K]) Insert(element E, key K) base.Handle {
	newNode := &leftistNode[E, K]{element: element, key: key, dist: 1}
	heap.min = heap.min.merge(newNode)
	heap.size++
	return newNode
}

// ExtractMin extracts the minimum element (with key/priority).
func (heap *LeftistHeap[E, K]) ExtractMin() (E, K) {
	if heap.IsEmpty() {
		panic("leftist heap empty")
	}

	minElement := heap.min.element
	minKey := heap.min.key

	// Merge the left and right subtrees of the root
	heap.min = heap.min.left.merge(heap.min.right)

	// Update parent pointer of the new root
	if heap.min != nil {
		heap.min.up = nil
	}

	heap.size--
	return minElement, minKey
}

// Remove removes an element (by handle) from the heap,
// splitting it in three parts, fixing one of them and
// merging them to one tree again.
func (heap *LeftistHeap[E, K]) Remove(handle base.Handle) {
	node, ok := handle.(*leftistNode[E, K])
	if !ok {
		panic("bad handle type")
	}

	// Case 1: Removing the root
	if node == heap.min {
		heap.min = heap.min.left.merge(heap.min.right)
		if heap.min != nil {
			heap.min.up = nil
		}
		heap.size--
		return
	}

	// Case 2: Removing an internal node

	// Get parent and children
	parent := node.up
	leftChild := node.left
	rightChild := node.right

	// Merge children
	merged := leftChild.merge(rightChild)

	// Connect merged subtree to parent
	if merged != nil {
		merged.up = parent
	}

	// Replace node in parent
	if parent.left == node {
		parent.left = merged
	} else {
		parent.right = merged
	}

	// Fix dist values upward
	current := parent
	for current != nil {
		// Recalculate dist and enforce leftist property
		if current.left == nil {
			current.dist = 1
		} else if current.right == nil {
			current.dist = 1
			// If left exists but right is nil, ensure left is the longer path
			if current.left != nil {
				current.left, current.right = current.right, current.left
			}
		} else {
			// Both children exist, ensure leftist property
			if current.left.dist < current.right.dist {
				current.left, current.right = current.right, current.left
			}
			current.dist = current.right.dist + 1
		}

		current = current.up
	}

	heap.size--
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

	// Store the element
	element := node.element

	// Remove the node and reinsert with new key
	heap.Remove(handle)
	heap.Insert(element, key)
}
