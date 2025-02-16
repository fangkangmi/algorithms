package main

import (
	"fmt"
)

// Graph structure using adjacency list representation
type Graph struct {
	adjList map[int][]int // Adjacency list to store the graph
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

// AddEdge adds an undirected edge between two nodes
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u) // For directed graph, remove this line
}

// BFS function to traverse all nodes using Breadth-First Search
func (g *Graph) BFS(startNode int) {
	visited := make(map[int]bool)
	queue := []int{startNode}

	for len(queue) != 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			node := queue[0]
			if visited[node] {
				queue = queue[1:]
				continue
			}
			visited[node] = true
			fmt.Println(node)
			queue = queue[1:]

			queue = append(queue, g.adjList[node]...)
		}
	}

	// What to take:
	// 1. for len(queue)>0 , dont have !=0

	// The answer

	// visited := make(map[int]bool)
	// queue := []int{startNode}

	// for len(queue) > 0 {
	// 	node := queue[0]
	// 	queue = queue[1:]

	// 	if !visited[node] {
	// 		visited[node] = true
	// 		fmt.Printf("Visited node: %d\n", node)

	// 		for _, neighbor := range g.adjList[node] {
	// 			if !visited[neighbor] {
	// 				queue = append(queue, neighbor)
	// 			}
	// 		}
	// 	}
	// }
}

// TODO: Implement your own traversal logic here if needed
func main() {
	graph := NewGraph()

	// Adding edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// Perform BFS starting from node 0
	fmt.Println("\nBFS Traversal:")
	graph.BFS(0)
}
