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

// DFS function to traverse all paths from startNode to endNode
func (g *Graph) DFSAllPaths(startNode, endNode int, path []int, visited map[int]bool, allPaths *[][]int) {
	// Mark the current node as visited and add it to the current path
	visited[startNode] = true
	path = append(path, startNode)

	// If we've reached the endNode, add the current path to allPaths
	if startNode == endNode {
		// Make a copy of the current path to store in allPaths
		tempPath := make([]int, len(path))
		copy(tempPath, path)
		*allPaths = append(*allPaths, tempPath)
	} else {
		// Otherwise, explore all neighbors
		for _, neighbor := range g.adjList[startNode] {
			if !visited[neighbor] {
				g.DFSAllPaths(neighbor, endNode, path, visited, allPaths)
			}
		}
	}

	// Backtrack: Remove the current node from the path and mark it as unvisited
	path = path[:len(path)-1]
	visited[startNode] = false
}

// Function to find all paths from startNode to endNode
func (g *Graph) FindAllPaths(startNode, endNode int) [][]int {
	var allPaths [][]int
	path := []int{}
	visited := make(map[int]bool)

	// Call the recursive DFS function to find all paths
	g.DFSAllPaths(startNode, endNode, path, visited, &allPaths)

	if len(allPaths) == 0 {
		return [][]int{}
	}

	return allPaths
}

func main() {
	graph := NewGraph()

	// Adding edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// Find all paths from node 0 to node 4
	startNode := 0
	endNode := 4
	allPaths := graph.FindAllPaths(startNode, endNode)

	// Print all paths
	fmt.Printf("All paths from %d to %d:\n", startNode, endNode)
	for _, path := range allPaths {
		fmt.Println(path)
	}
}
