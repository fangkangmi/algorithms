package main

import (
	"fmt"
)

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) PopMin() int {
	if len(h.arr) == 0 {
		return -1
	}
	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyUp(index int) {
	for h.arr[parent(index)] > h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] < h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.arr[index] > h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func print_minheap() {
	h := &MinHeap{}
	h.Insert(3)
	h.Insert(1)
	h.Insert(6)
	h.Insert(5)
	h.Insert(2)
	h.Insert(4)

	fmt.Println(h.PopMin()) // 1
	fmt.Println(h.PopMin()) // 2
	fmt.Println(h.PopMin()) // 3
	fmt.Println(h.PopMin()) // 4
	fmt.Println(h.PopMin()) // 5
	fmt.Println(h.PopMin()) // 6
}
