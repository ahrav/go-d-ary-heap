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
		operations []heapOperation[int]
	}{
		{
			name: "MinHeap with d=2",
			d:    2,
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
			name: "Single element",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 10},
				{operation: "peek", want: 10},
				{operation: "pop", want: 10, wantHeap: []int{}},
			},
		},
		{
			name: "All elements same",
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
			name: "Ascending order push",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 1},
				{operation: "push", value: 2},
				{operation: "push", value: 3},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{2, 3}},
			},
		},
		{
			name: "Descending order push",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "push", value: 3},
				{operation: "push", value: 2},
				{operation: "push", value: 1},
				{operation: "peek", want: 1},
				{operation: "pop", want: 1, wantHeap: []int{2, 3}},
			},
		},
		{
			name: "Higher branching factor",
			d:    4,
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
			name: "Pop from empty heap",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "pop", want: 0, wantHeap: []int{}}, // Assuming default zero value is returned for empty heap
			},
		},
		{
			name: "Peek into empty heap",
			d:    2,
			operations: []heapOperation[int]{
				{operation: "peek", want: 0}, // Assuming default zero value is returned for empty heap
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			heap := NewHeap[int](tt.d, func(a, b int) bool { return a < b })

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
					for i, wantVal := range op.wantHeap {
						if i >= heap.heapSize || heap.data[i] != wantVal {
							t.Errorf("After %s, heap[%d] = %v, want %v", op.operation, i, heap.data[i], wantVal)
							break
						}
					}
					assert.Equal(t, len(op.wantHeap), heap.heapSize, "heap size not as expected after %s", op.operation)
				}
			}
		})
	}
}
