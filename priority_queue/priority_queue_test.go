package main

import (
	"testing"
)

func TestSimpleMinPQ(t *testing.T) {
	pq := SimpleMinPQ{
		heap: make([]int, 10),
		size: 0,
	}

	// Test case 1: Push elements in random order and pop them
	pq.push(10)
	pq.push(4)
	pq.push(7)
	pq.push(1)
	pq.push(3)
	pq.push(9)
	pq.push(2)
	pq.push(8)
	pq.push(5)
	pq.push(6)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range expected {
		if res := pq.pop(); res != v {
			t.Errorf("Expected %d, but got %d", v, res)
		}
	}

	// Test case 2: Push elements in descending order and pop them
	pq.push(5)
	pq.push(4)
	pq.push(3)
	pq.push(2)
	pq.push(1)

	expected = []int{1, 2, 3, 4, 5}
	for _, v := range expected {
		if res := pq.pop(); res != v {
			t.Errorf("Expected %d, but got %d", v, res)
		}
	}

	// Test case 3: Push elements in ascending order and pop them
	pq.push(1)
	pq.push(2)
	pq.push(3)
	pq.push(4)
	pq.push(5)

	expected = []int{1, 2, 3, 4, 5}
	for _, v := range expected {
		if res := pq.pop(); res != v {
			t.Errorf("Expected %d, but got %d", v, res)
		}
	}

	// Test case 4: Push and pop elements alternately
	pq.push(3)
	if res := pq.pop(); res != 3 {
		t.Errorf("Expected 3, but got %d", res)
	}
	pq.push(2)
	if res := pq.pop(); res != 2 {
		t.Errorf("Expected 2, but got %d", res)
	}
	pq.push(1)
	if res := pq.pop(); res != 1 {
		t.Errorf("Expected 1, but got %d", res)
	}
	pq.push(5)
	if res := pq.pop(); res != 5 {
		t.Errorf("Expected 5, but got %d", res)
	}
	pq.push(4)
	if res := pq.pop(); res != 4 {
		t.Errorf("Expected 4, but got %d", res)
	}
}
