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

// DFS function to traverse all nodes using Depth-First Search
func (g *Graph) DFS(node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("Visited node: %d\n", node)

	for _, neighbor := range g.adjList[node] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
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

	// Perform DFS starting from node 0
	fmt.Println("DFS Traversal:")
	visitedDFS := make(map[int]bool)
	graph.DFS(0, visitedDFS)

}
