package main

import (
	"fmt"
)

type SimpleMinPQ struct {
	// 底层使用数组实现二叉堆
	heap []int

	// 堆中元素的数量
	size int
}

// 父节点的索引
func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

// 左子节点的索引
func (pq *SimpleMinPQ) left(node int) int {
	return node*2 + 1
}

// 右子节点的索引
func (pq *SimpleMinPQ) right(node int) int {
	return node*2 + 2
}

// 交换数组的两个元素
func (pq *SimpleMinPQ) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

// 查，返回堆顶元素，时间复杂度 O(1)
func (pq *SimpleMinPQ) peek() int {
	return pq.heap[0]
}

// 增，向堆中插入一个元素，时间复杂度 O(logN)
func (pq *SimpleMinPQ) push(x int) {
	pq.heap = append(pq.heap, x) // Append the new element
	pq.swim(pq.size)             // Swim to restore heap property
	pq.size++
}

// 删，删除堆顶元素，时间复杂度 O(logN)
func (pq *SimpleMinPQ) pop() int {
	res := pq.heap[0]
	pq.swap(0, pq.size-1)
	pq.heap = pq.heap[:pq.size-1]
	pq.size--
	pq.sink(0)
	return res
}

// 上浮操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.heap[node] < pq.heap[pq.parent(node)] {
		pq.swap(pq.parent(node), node)
		node = pq.parent(node)
	}
}

// 下沉操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) sink(node int) {
	// if the node has left and right nodes
	for pq.left(node) < pq.size && pq.right(node) < pq.size {

		// if the node is greater than left or right, we swap it
		if pq.heap[node] > pq.heap[pq.left(node)] || pq.heap[node] > pq.heap[pq.right(node)] {

			// if left less than right, we swap the left to node
			if pq.heap[pq.left(node)] < pq.heap[pq.right(node)] {
				// we first swap the value
				pq.swap(pq.left(node), node)

				// Then we swap the index
				node = pq.left(node)
			} else {
				pq.swap(pq.right(node), node)
				node = pq.right(node)
			}
		}
	}

	if pq.left(node) < pq.size {
		if pq.heap[pq.left(node)] < pq.heap[node] {
			pq.swap(pq.left(node), node)
		}
	}
}

// 下沉操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) sink_recursion(node int) {

	// base case
	if pq.left(node) >= pq.size {
		return
	}

	min := node
	left := pq.left(node)
	right := pq.right(node)

	if left < pq.size && pq.heap[left] < pq.heap[min] {
		min = left
	}

	if right < pq.size && pq.heap[right] < pq.heap[min] {
		min = right
	}

	if min == node {
		return
	}

	pq.swap(node, min)
	pq.sink(min)

}

func main() {
	pq := SimpleMinPQ{
		heap: []int{},
		size: 0,
	}

	pq.push(3)
	pq.push(2)
	pq.push(1)
	pq.push(5)
	pq.push(4)
	fmt.Println("original heap")
	fmt.Println(pq.heap)

	fmt.Println("start pop")
	pq.pop()
	fmt.Println(pq.heap)
	pq.pop()
	fmt.Println(pq.heap)
	pq.pop()
	fmt.Println(pq.heap)
	pq.pop()
	fmt.Println(pq.heap)
	pq.pop()
	fmt.Println(pq.heap)
}
