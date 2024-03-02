# D-Ary Heap Implementation in Go

## Introduction

This package provides a generic implementation of a d-ary heap in Go, suitable for any ordered type.
A d-ary heap is a generalization of a binary heap where each node has `d` children instead of two.
This variation allows for a more shallow heap, potentially optimizing operations like decrease-key,
which benefits from a shorter path from any given node to the root.

## Why Use a D-Ary Heap?

The d-ary heap is particularly useful in scenarios where heap operations are a bottleneck in performance.
Due to its more shallow structure compared to a binary heap, operations that move elements between the root
and the leaf nodes (like insertions and deletions) can be faster, as they generally involve fewer steps.
This makes d-ary heaps an excellent choice for priority queues, especially when managing large datasets or
when fine-tuning performance is necessary.

## Installation

To use this package, simply import it into your Go project:

```go
import "github.com/ahrav/go-d-ary-heap"
```

## Usage

Below are examples demonstrating how to use this d-ary heap implementation with different types and operations.

### Creating a New Heap

```go
package main

import (
    "fmt"
    "github.com/ahrav/go-d-ary-heap"
)

func main() {
    // Create a min-heap for integers with a branching factor of 3.
    minHeap := heap.NewHeap[int](3, func(a, b int) bool { return a < b })

    // Create a max-heap for integers with a branching factor of 4.
    maxHeap := heap.NewHeap[int](4, func(a, b int) bool { return a > b })
}
```

### Adding Elements

```go
minHeap.Push(10)
minHeap.Push(5)
minHeap.Push(15)

maxHeap.Push(10)
maxHeap.Push(5)
maxHeap.Push(15)
```

### Removing the Extremal Element

```go
fmt.Println(minHeap.Pop()) // Outputs: 5
fmt.Println(maxHeap.Pop()) // Outputs: 15
```

### Peeking at the Extremal Element

```go
fmt.Println(minHeap.Peek()) // Assuming more elements were added, outputs the smallest
fmt.Println(maxHeap.Peek()) // Assuming more elements were added, outputs the largest
```

## Contributing

Contributions to improve the d-ary heap implementation are welcome.
Please feel free to submit pull requests or open issues to discuss potential improvements or report bugs.
