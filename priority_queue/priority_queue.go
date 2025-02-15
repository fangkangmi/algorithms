package main

import "fmt"

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
	// 把新元素追加到最后
	pq.heap[pq.size] = x
	// 然后上浮到正确位置
	pq.swim(pq.size)
	pq.size++
}

// 删，删除堆顶元素，时间复杂度 O(logN)
func (pq *SimpleMinPQ) pop() int {
	res := pq.heap[0]
	// 把堆底元素放到堆顶
	pq.heap[0] = pq.heap[pq.size-1]
	pq.size--
	// move the element at the root (index 0) down to its correct position
	pq.sink(0)
	return res
}

// 上浮操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.heap[pq.parent(node)] > pq.heap[node] {
		pq.swap(pq.parent(node), node)
		node = pq.parent(node)
	}
}

// 下沉操作，时间复杂度是树高 O(logN)
func (pq *SimpleMinPQ) sink(node int) {
	for pq.left(node) < pq.size || pq.right(node) < pq.size {
		// 比较自己和左右子节点，看看谁最小
		min := node
		if pq.left(node) < pq.size && pq.heap[pq.left(node)] < pq.heap[min] {
			min = pq.left(node)
		}
		if pq.right(node) < pq.size && pq.heap[pq.right(node)] < pq.heap[min] {
			min = pq.right(node)
		}
		if min == node {
			break
		}
		// 如果左右子节点中有比自己小的，就交换
		pq.swap(node, min)
		node = min
	}
}

func main() {
	fmt.Println("testing")
}
