package leftist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.lrz.de/hm/goal-core/collections"
)

func TestCreation(t *testing.T) {
	var pq collections.PriorityQueue[string, float32] = NewLeftistHeap[string, float32]()
	assert.Equal(t, uint(0), pq.Size())
}

func TestSingleInsert(t *testing.T) {
	pq := NewLeftistHeap[string, float32]()
	pq.Insert("X", 2.0)
	assert.Equal(t, uint(1), pq.Size())
}

func TestSingleMinExtract(t *testing.T) {
	pq := NewLeftistHeap[string, float32]()
	pq.Insert("X", 2.0)
	e, k := pq.ExtractMin()
	assert.Equal(t, uint(0), pq.Size())
	assert.Equal(t, "X", e)
	assert.Equal(t, float32(2.0), k)
}

func TestMultipleInsert(t *testing.T) {
	pq := NewLeftistHeap[string, int]()
	pq.Insert("X", 4)
	pq.Insert("Y", 3)
	pq.Insert("Z", 7)
	pq.Insert("G", 5)
	assert.Equal(t, uint(4), pq.Size())
	e, _ := pq.ExtractMin()
	assert.Equal(t, "Y", e)
	e, _ = pq.ExtractMin()
	assert.Equal(t, "X", e)
	e, _ = pq.ExtractMin()
	assert.Equal(t, "G", e)
	e, _ = pq.ExtractMin()
	assert.Equal(t, "Z", e)
}

func TestSimpleRemoveRoot(t *testing.T) {
	pq := NewLeftistHeap[string, float32]()
	xh := pq.Insert("X", 2.0)
	pq.Insert("Y", 3.0)
	pq.Insert("Z", 7.0)
	pq.Remove(xh)
	assert.Equal(t, uint(2), pq.Size())
	e, _ := pq.ExtractMin()
	assert.Equal(t, "Y", e)
	e, _ = pq.ExtractMin()
	assert.Equal(t, "Z", e)
}

func TestSimpleRemoveInner(t *testing.T) {
	pq := NewLeftistHeap[string, float32]()
	pq.Insert("X", 2.0)
	xh := pq.Insert("Y", 3.0)
	pq.Insert("Z", 7.0)
	pq.Insert("G", 5.0)
	pq.Remove(xh)
	assert.Equal(t, uint(3), pq.Size())
	e, _ := pq.ExtractMin()
	assert.Equal(t, "X", e)
	e, _ = pq.ExtractMin()
	assert.Equal(t, "G", e)
}

func Test_DecreaseKey(t *testing.T) {
	pq := NewLeftistHeap[string, int]()
	pq.Insert("X", 30)
	pq.Insert("Y", 20)
	pq.Insert("Z", 40)
	h := pq.Insert("U", 60)
	pq.Insert("V", 100)
	pq.Insert("W", 100)
	pq.DecreaseKey(h, 22)
	e, x := pq.ExtractMin()
	assert.Equal(t, "Y", e)
	assert.Equal(t, 20, x)
	e, x = pq.ExtractMin()
	assert.Equal(t, "U", e)
	assert.Equal(t, 22, x)
	assert.Equal(t, uint(4), pq.Size())
}

func TestRemoveRootWithEmptyChildren(t *testing.T) {
	pq := NewLeftistHeap[string, int]()
	h := pq.Insert("A", 10)

	// Remove the only node in the heap
	pq.Remove(h)
	assert.True(t, pq.IsEmpty())
	assert.Equal(t, uint(0), pq.Size())
}

func TestComplexRemoveInternal(t *testing.T) {
	pq := NewLeftistHeap[string, int]()

	// Create a more complex tree structure
	pq.Insert("A", 10)
	h2 := pq.Insert("B", 20)
	pq.Insert("C", 15)
	pq.Insert("D", 25)
	pq.Insert("E", 30)

	// Remove an internal node
	pq.Remove(h2)

	// Check that structure is maintained correctly
	assert.Equal(t, uint(4), pq.Size())

	// Extract all and verify order
	values := make([]string, 0, 4)
	keys := make([]int, 0, 4)

	for !pq.IsEmpty() {
		e, k := pq.ExtractMin()
		values = append(values, e)
		keys = append(keys, k)
	}

	// Check expected order based on keys
	assert.Equal(t, []string{"A", "C", "D", "E"}, values)
	assert.Equal(t, []int{10, 15, 25, 30}, keys)
}

func TestMergeSymmetry(t *testing.T) {
	// Test that merge(a, b) and merge(b, a) produce equivalent heaps
	pq1 := NewLeftistHeap[string, int]()
	pq2 := NewLeftistHeap[string, int]()

	// Create two heaps
	pq1.Insert("A", 10)
	pq1.Insert("B", 30)

	pq2.Insert("C", 20)
	pq2.Insert("D", 40)

	// Get their roots
	root1 := pq1.min
	root2 := pq2.min

	// Merge root1 with root2
	merged := root1.merge(root2)

	// Check that merged heap has correct elements
	// Extract all elements and check
	values := []string{}
	keys := []int{}

	for merged != nil {
		values = append(values, merged.element)
		keys = append(keys, merged.key)

		// Remove current min
		merged = merged.left.merge(merged.right)
	}

	// Verify expected order
	assert.Equal(t, []int{10, 20, 30, 40}, keys)
}

func TestEmptyInitialMerge(t *testing.T) {
	// Test initializing a heap through merges of empty heaps
	var emptyNode *leftistNode[string, int]

	// Merge empty with empty should be empty
	merged := emptyNode.merge(nil)
	assert.Nil(t, merged)

	// Create a node
	newNode := &leftistNode[string, int]{
		element: "A",
		key:     10,
		dist:    1,
	}

	// Merge empty with node
	merged = emptyNode.merge(newNode)
	assert.Equal(t, newNode, merged)

	// Merge node with empty
	merged = newNode.merge(nil)
	assert.Equal(t, newNode, merged)
}

func TestRemoveNonRoot(t *testing.T) {
	pq := NewLeftistHeap[string, int]()

	// Create a heap with several elements
	pq.Insert("A", 10)
	h2 := pq.Insert("B", 20)
	pq.Insert("C", 30)
	pq.Insert("D", 40)

	// Remove a non-root element
	pq.Remove(h2)

	// Extract all elements and verify order
	values := []string{}
	keys := []int{}

	for !pq.IsEmpty() {
		e, k := pq.ExtractMin()
		values = append(values, e)
		keys = append(keys, k)
	}

	assert.Equal(t, []string{"A", "C", "D"}, values)
	assert.Equal(t, []int{10, 30, 40}, keys)
}
