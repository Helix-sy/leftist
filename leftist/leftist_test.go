package leftist

import (
	"github.com/stretchr/testify/assert"
	"gitlab.lrz.de/hm/goal-core/collections"
	"testing"
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
