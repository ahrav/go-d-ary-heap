package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestHeapOperations(t *testing.T) {
	type heapOperation[T constraints.Ordered] struct {
		operation string // "push", "pop", or "peek"
		value     T      // value to push (ignored for pop and peek)
		want      T      // expected result for peek or pop
		wantHeap  []T    // expected heap state after the operation
	}

	tests := []struct {
		name       string
		d          int
		isMinHeap  bool
		operations []heapOperation[int]
	}{
		{
			name:      "MinHeap with d=2",
			d:         2,
			isMinHeap: true,
			operations: []heapOperation[int]{
				{operation: "push", value: 5},
				{operation: "push", value: 3},
				{operation: "peek", want: 3},
				{operation: "push", value: 4},
				{operation: "pop", want: 3, wantHeap: []int{4, 5}},
				{operation: "push", value: 1},
				{operation: "pop", want: 1, wantHeap: []int{4, 5}},
				{operation: "pop", want: 4, wantHeap: []int{5}},
				{operation: "pop", want: 5, wantHeap: []int{}},
			},
		},
		{
			name:      "MinHeap Single element",
			d:         2,
			isMinHeap: true,
			operations: []heapOperation[int]{
				{operation: "push", value: 10},
				{operation: "peek", want: 10},
				{operation: "pop", want: 10, wantHeap: []int{}},
			},
		},
		{
			name: "MinHeap All elements same",
			d:    3,
			operations: []heapOperation[int]{
				{operation: "push", value: 1},
				{operation: "push", value: 1},
				{operation: "push", value: 1},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{1, 1}},
				{operation: "pop", want: 1, wantHeap: []int{1}},
			},
		},
		{
			name:      "MinHeap Ascending order push",
			d:         2,
			isMinHeap: true,
			operations: []heapOperation[int]{
				{operation: "push", value: 1},
				{operation: "push", value: 2},
				{operation: "push", value: 3},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{2, 3}},
			},
		},
		{
			name:      "MinHeap Descending order push",
			d:         2,
			isMinHeap: true,
			operations: []heapOperation[int]{
				{operation: "push", value: 3},
				{operation: "push", value: 2},
				{operation: "push", value: 1},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{2, 3}},
			},
		},
		{
			name:      "MinHeap Higher branching factor",
			d:         4,
			isMinHeap: true,
			operations: []heapOperation[int]{
				{operation: "push", value: 5},
				{operation: "push", value: 3},
				{operation: "push", value: 4},
				{operation: "push", value: 1},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{3, 5, 4}},
				{operation: "pop", want: 3, wantHeap: []int{4, 5}},
				{operation: "push", value: 4},
				{operation: "pop", want: 4, wantHeap: []int{4, 5}},
				{operation: "pop", want: 4, wantHeap: []int{5}},
				{operation: "pop", want: 5, wantHeap: []int{}},
			},
		},
		{
			name:      "MaxHeap with d=2",
			d:         2,
			isMinHeap: false,
			operations: []heapOperation[int]{
				{operation: "push", value: 5},
				{operation: "push", value: 3},
				{operation: "peek", want: 5},
				{operation: "push", value: 4},
				{operation: "pop", want: 5, wantHeap: []int{4, 3}},
				{operation: "push", value: 1},
				{operation: "pop", want: 4, wantHeap: []int{3, 1}},
				{operation: "pop", want: 3, wantHeap: []int{1}},
				{operation: "pop", want: 1, wantHeap: []int{}},
			},
		},
		{
			name: "MaxHeap Single element",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 10},
				{operation: "peek", want: 10},
				{operation: "pop", want: 10, wantHeap: []int{}},
			},
		},
		{
			name: "MaxHeap All elements same",
			d:    3,
			operations: []heapOperation[int]{
				{operation: "push", value: 1},
				{operation: "push", value: 1},
				{operation: "push", value: 1},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{1, 1}},
				{operation: "pop", want: 1, wantHeap: []int{1}},
			},
		},
		{
			name: "MaxHeap Ascending order push",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 1},
				{operation: "push", value: 2},
				{operation: "push", value: 3},
				{operation: "peek", want: 3},
				{operation: "pop", want: 3, wantHeap: []int{2, 1}},
			},
		},
		{
			name: "MaxHeap Descending order push",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 3},
				{operation: "push", value: 2},
				{operation: "push", value: 1},
				{operation: "peek", want: 3},
				{operation: "pop", want: 3, wantHeap: []int{2, 1}},
			},
		},
		{
			name: "MaxHeap Higher branching factor",
			d:    4,
			operations: []heapOperation[int]{
				{operation: "push", value: 5},
				{operation: "push", value: 3},
				{operation: "push", value: 4},
				{operation: "push", value: 1},
				{operation: "peek", want: 5},
				{operation: "pop", want: 5, wantHeap: []int{4, 3, 1}},
				{operation: "pop", want: 4, wantHeap: []int{3, 1}},
				{operation: "push", value: 4},
				{operation: "pop", want: 4, wantHeap: []int{3, 1}},
				{operation: "pop", want: 3, wantHeap: []int{1}},
				{operation: "pop", want: 1, wantHeap: []int{}},
			},
		},
		{
			name: "MaxHeap Pop from empty heap",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "pop", want: 0, wantHeap: []int{}},
			},
		},
		{
			name: "MaxHeap Peek into empty heap",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "peek", want: 0},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var heap *Heap[int]
			if tt.isMinHeap {
				heap = NewHeap[int](tt.d, func(a, b int) bool { return a < b })
			} else {
				heap = NewHeap[int](tt.d, func(a, b int) bool { return a > b })
			}

			for _, op := range tt.operations {
				switch op.operation {
				case "push":
					heap.Push(op.value)
				case "pop":
					assert.Equal(t, op.want, heap.Pop(), "Pop() returned wrong value")
				case "peek":
					assert.Equal(t, op.want, heap.Peek(), "Peek() returned wrong value")
				default:
					t.Fatalf("unknown operation: %s", op.operation)
				}

				if op.wantHeap != nil {
					if len(op.wantHeap) > 0 {
						for i, wantVal := range op.wantHeap {
							if i >= heap.heapSize || heap.data[i] != wantVal {
								t.Errorf("After %s, heap[%d] = %v, want %v", op.operation, i, heap.data[i], wantVal)
								break
							}
						}
					}
					assert.Equal(t, len(op.wantHeap), heap.heapSize, "heap size not as expected after %s", op.operation)
				}
			}
		})
	}
}

func TestHeapContains(t *testing.T) {
	heap := NewHeap[int](2, func(a, b int) bool { return a < b })
	heap.Push(5)
	heap.Push(3)
	heap.Push(4)
	heap.Push(1)
	heap.Push(1)

	assert.True(t, heap.Contains(5), "Contains(5) returned false, want true")
	assert.True(t, heap.Contains(3), "Contains(3) returned false, want true")
	assert.False(t, heap.Contains(2), "Contains(2) returned true, want false")

	// Ensure duplicates are handled correctly.
	heap.Pop()
	assert.True(t, heap.Contains(1), "Contains(5) returned true, want false")

	heap.Pop()
	assert.False(t, heap.Contains(1), "Contains(1) returned true, want false")
}

func TestHeapGet(t *testing.T) {
	heap := NewHeap[int](2, func(a, b int) bool { return a < b })
	heap.Push(5)
	heap.Push(3)
	heap.Push(4)
	heap.Push(1)
	heap.Push(1)

	val, ok := heap.Get(5)
	assert.True(t, ok, "Get(5) returned false, want true")
	assert.Equal(t, 5, val, "Get(5) returned %d, want 5", val)

	val, ok = heap.Get(3)
	assert.True(t, ok, "Get(3) returned false, want true")
	assert.Equal(t, 3, val, "Get(3) returned %d, want 3", val)

	val, ok = heap.Get(2)
	assert.False(t, ok, "Get(2) returned true, want false")
	assert.Zero(t, val, "Get(2) returned %d, want 0", val)

	// Ensure duplicates are handled correctly.
	heap.Pop()
	val, ok = heap.Get(1)
	assert.True(t, ok, "Get(1) returned false, want true")
	assert.Equal(t, 1, val, "Get(1) returned %d, want 1", val)

	heap.Pop()
	val, ok = heap.Get(1)
	assert.False(t, ok, "Get(1) returned true, want false")
	assert.Zero(t, val, "Get(1) returned %d, want 0", val)
}
