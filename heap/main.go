package main

import (
	"container/heap"
	"fmt"
	"main/internal/intHeap"
)

func main() {
	h := &intHeap.IntHeap{1, 2, 3}
	heap.Init(h)
	heap.Push(h, 1)
	heap.Push(h, 5)

	for len(*h) > 0 {
		fmt.Println(heap.Pop(h))
	}
}
