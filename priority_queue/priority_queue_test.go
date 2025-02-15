package main

import (
	"testing"
)

func TestSink(t *testing.T) {
	tests := []struct {
		initialHeap  []int
		expectedHeap []int
	}{
		{[]int{3, 2, 1, 5, 4}, []int{1, 2, 3, 5, 4}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 4, 3, 6, 5, 8, 7, 9, 10}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 5, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 6, 5, 4, 7, 9, 8}},
	}

	for _, test := range tests {
		pq := SimpleMinPQ{
			heap: test.initialHeap,
			size: len(test.initialHeap),
		}
		pq.sink(0)
		for i, v := range pq.heap {
			if v != test.expectedHeap[i] {
				t.Errorf("Test failed for initial heap %v. Expected %v, got %v", test.initialHeap, test.expectedHeap, pq.heap)
				break
			}
		}
	}
}
