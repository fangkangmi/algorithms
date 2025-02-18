package main

import "fmt"

// Node represents a node in the doubly-linked list.
type Node struct {
	key, value interface{}
	prev, next *Node
}

// LinkedHashMap represents the LinkedHashMap data structure.
type LinkedHashMap struct {
	// Map to store key-node mappings
	mapData map[interface{}]*Node

	// Sentinel nodes for the doubly-linked list
	head *Node
	tail *Node

	// Size of the LinkedHashMap
	size int
}

// NewLinkedHashMap creates and initializes a new LinkedHashMap.
func NewLinkedHashMap() *LinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LinkedHashMap{
		mapData: make(map[interface{}]*Node),
		head:    head,
		tail:    tail,
		size:    0,
	}
}

// Add inserts a key-value pair into the LinkedHashMap.
func (lhm *LinkedHashMap) Add(key, value interface{}) {
	if node, exists := lhm.mapData[key]; exists {
		// Update the existing node's value and move it to the end (most recently used)
		node.value = value
		lhm.moveToEnd(node)
	} else {
		// Create a new node and add it to the map and the linked list
		newNode := &Node{key: key, value: value}
		lhm.mapData[key] = newNode
		lhm.appendNode(newNode)
		lhm.size++
	}
}

// Get retrieves the value associated with the given key and moves the node to the end (most recently used).
func (lhm *LinkedHashMap) Get(key interface{}) (interface{}, bool) {
	if node, exists := lhm.mapData[key]; exists {
		lhm.moveToEnd(node)
		return node.value, true
	}
	return nil, false
}

// Remove deletes the key-value pair associated with the given key.
func (lhm *LinkedHashMap) Remove(key interface{}) {
	if node, exists := lhm.mapData[key]; exists {
		lhm.removeNode(node)
		delete(lhm.mapData, key)
		lhm.size--
	}
}

// Keys returns the keys in the order they were inserted.
func (lhm *LinkedHashMap) Keys() []interface{} {
	keys := make([]interface{}, 0, lhm.size)
	for current := lhm.head.next; current != lhm.tail; current = current.next {
		keys = append(keys, current.key)
	}
	return keys
}

// Size returns the number of key-value pairs in the LinkedHashMap.
func (lhm *LinkedHashMap) Size() int {
	return lhm.size
}

// Helper function to append a node to the end of the list.
func (lhm *LinkedHashMap) appendNode(node *Node) {
	last := lhm.tail.prev
	last.next = node
	node.prev = last
	node.next = lhm.tail
	lhm.tail.prev = node
}

// Helper function to remove a node from the list.
func (lhm *LinkedHashMap) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// Helper function to move a node to the end of the list.
func (lhm *LinkedHashMap) moveToEnd(node *Node) {
	lhm.removeNode(node)
	lhm.appendNode(node)
}

// Example usage
func main() {
	lhm := NewLinkedHashMap()

	// Adding elements
	lhm.Add("one", 1)
	lhm.Add("two", 2)
	lhm.Add("three", 3)

	// Getting elements
	if value, ok := lhm.Get("two"); ok {
		fmt.Println("Value for 'two':", value)
	}

	// Removing an element
	lhm.Remove("one")

	// Printing keys in insertion order
	fmt.Println("Keys in insertion order:", lhm.Keys())

	// Adding another element
	lhm.Add("four", 4)

	// Printing keys again
	fmt.Println("Keys in insertion order after adding 'four':", lhm.Keys())
}

// queue := []int{0}
// var res [][]int
// path:= []int{0}

// last_node := len(graph)-1
// for len(queue)>0{

// 	sz := len(queue)
// 	// Here is to iterate the size of next level
// 	// If we create all the path from 0 to any points. The space COmplx will be O(V+E), too big\
// 	// So we only save: When there's no neighbor.
// 	// But we still need to save the state, before reaching 'no neighbor'
// 	// Luckily, the 'no neighbor' point is the one and the only solution
// 	// So wee need to save all the path to the end

// 	// What the pain pioint here is to: don't save every path, but extend existing paths
// 	// And make sure it is extended the correct path.
// 	// So let us for every loop, we need to have a path for "How we have been here"

// 	for i := 0 ; i < sz; i++{
// 		node := queue[0]
// 		queue = queue[:1]
// 		// Iterate that specific node's neighbor,
// 		// we should have a path here, and extend them in the for loop
// 		// Also in for loop we need to create some paths
// 		for _, neighbor := range graph[node]{
// 			queue = append(queue,neighbor)
// 			all_path = append(all_path, )

// 		}
// 	}
// 	path = append(path,)
// }
