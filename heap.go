// Package heap provides operations for a generic d-ary heap. Unlike the standard
// library's heap package, which requires types to implement the heap.Interface,
// this package offers a concrete implementation of a d-ary heap that works with
// any ordered type, as defined by the constraints.Ordered constraint from
// golang.org/x/exp/constraints.
//
// A d-ary heap is a variation of the binary heap where each node can have up to
// d children instead of just two. This allows for a more shallow heap for the
// same number of elements, potentially optimizing certain operations like
// decrease-key. The heap maintains the property that each node is ordered
// according to a provided less function, ensuring that the root of the heap
// always contains the extremal element (minimum or maximum).
//
// The Heap struct in this package encapsulates the d-ary heap's state, including
// the heap's elements, its branching factor (d), and a custom less function to
// determine the order of elements. This implementation allows for a flexible and
// generic heap that can handle any ordered type without requiring additional
// methods on the type itself.
//
// Basic operations provided include:
// - NewHeap: to initialize a new d-ary heap with a specified branching factor and ordering function.
// - Push: to add new elements to the heap while maintaining the heap property.
// - Pop: to remove and return the extremal element from the heap.
// - Peek: to return the extremal element without removing it.
// - Update: to change an element's value and then restore the heap property.
//
// This package is designed for use cases where a priority queue or any other
// application requires a dynamically ordered set of elements and can benefit
// from the efficiency of a heap, especially when the optimal branching factor
// may differ from the binary case.

package heap

import (
	"golang.org/x/exp/constraints"
)

// Heap struct represents a generic d-ary heap.
type Heap[T constraints.Ordered] struct {
	data     []T             // Underlying array to store the heap elements
	d        int             // Branching factor (number of children per node)
	heapSize int             // Current size of the heap
	lessFunc func(T, T) bool // Function to determine order
}

// Option is a type representing configurations for the heap
type Option[T constraints.Ordered] func(*Heap[T])

// WithCapacity is an option that sets the initial capacity of the heap
func WithCapacity[T constraints.Ordered](capacity int) Option[T] {
	return func(h *Heap[T]) {
		h.data = make([]T, capacity)
	}
}

// NewHeap creates a new d-ary heap with the specified branching factor.
func NewHeap[T constraints.Ordered](d int, less func(T, T) bool, options ...Option[T]) *Heap[T] {
	heap := &Heap[T]{
		d:        d,
		data:     make([]T, 1), // Start with an initial capacity of 1
		heapSize: 0,
		lessFunc: less,
	}

	for _, option := range options {
		option(heap)
	}

	return heap
}

// parent returns the index of the parent node for a given index.
func (h *Heap[T]) parent(i int) int {
	return (i - 1) / h.d
}

// child returns the index of the k-th child of a given index.
func (h *Heap[T]) child(i, k int) int {
	return h.d*i + k
}

// less determines if the element at index i is less than the element at index j.
func (h *Heap[T]) less(i, j int) bool {
	return h.data[i] < h.data[j]
}

// swap swaps the elements at indices i and j.
func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Peek returns the minimum element from the heap without removing it.
func (h *Heap[T]) Peek() T { return h.data[0] }

// Push adds a new element to the heap.
func (h *Heap[T]) Push(value T) {
	if h.heapSize == len(h.data) {
		h.data = append(h.data, value)
	} else {
		h.data[h.heapSize] = value
	}

	h.heapSize++
	h.up(h.heapSize - 1) // Restore heap property after insertion
}

// Pop removes and returns the minimum element from the heap.
func (h *Heap[T]) Pop() T {
	var minV T
	if h.heapSize == 0 {
		return minV
	}

	minV = h.data[0]
	h.swap(0, h.heapSize-1)
	h.heapSize--
	h.down(0) // Restore heap property after removal
	return minV
}

// up restores the heap property by bubbling an element up the tree.
func (h *Heap[T]) up(i int) {
	for i > 0 && h.less(i, h.parent(i)) {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

// down restores the heap property by moving an element down the tree.
func (h *Heap[T]) down(i int) {
	for {
		smallest := i // Assume the current node is the smallest
		for k := 1; k <= h.d && h.child(i, k) < h.heapSize; k++ {
			childIndex := h.child(i, k)
			if h.less(childIndex, smallest) {
				smallest = childIndex
			}
		}

		if smallest == i {
			break // Heap property is satisfied
		}
		h.swap(i, smallest)
		i = smallest
	}
}
