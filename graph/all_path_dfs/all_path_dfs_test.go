package main

import (
	"reflect"
	"testing"
)

// Test Case 1: Simple Graph
func TestSimpleGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)

	startNode := 0
	endNode := 5
	expected := [][]int{{0, 2, 5}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestSimpleGraph failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 2: Graph with Multiple Paths
func TestMultiplePaths(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(4, 6)

	startNode := 0
	endNode := 6
	expected := [][]int{{0, 1, 4, 6}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestMultiplePaths failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 3: Graph with Cycles
func TestGraphWithCycles(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(4, 6)

	startNode := 0
	endNode := 6
	expected := [][]int{{0, 1, 3, 6}, {0, 1, 4, 6}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestGraphWithCycles failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 4: Directed Graph
func TestDirectedGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)

	startNode := 0
	endNode := 5
	expected := [][]int{{0, 2, 5}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestDirectedGraph failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 5: Disconnected Graph
func TestDisconnectedGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(6, 7)

	startNode := 0
	endNode := 7
	expected := [][]int{}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestDisconnectedGraph failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 6: Graph with Self-Loops
func TestGraphWithSelfLoops(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 3)
	graph.AddEdge(5, 5)

	startNode := 0
	endNode := 5
	expected := [][]int{{0, 2, 5}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestGraphWithSelfLoops failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 7: Graph with Multiple Cycles
func TestGraphWithMultipleCycles(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(4, 6)
	graph.AddEdge(4, 7)
	graph.AddEdge(5, 7)

	startNode := 0
	endNode := 7
	expected := [][]int{{0, 1, 3, 6, 4, 7}, {0, 1, 4, 7}, {0, 2, 5, 7}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestGraphWithMultipleCycles failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 8: Large Graph
func TestLargeGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)
	graph.AddEdge(2, 6)
	graph.AddEdge(3, 7)
	graph.AddEdge(3, 8)

	startNode := 0
	endNode := 8
	expected := [][]int{{0, 3, 8}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestLargeGraph failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 9: Single Node Graph
func TestSingleNodeGraph(t *testing.T) {
	graph := NewGraph()

	startNode := 0
	endNode := 0
	expected := [][]int{{0}}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestSingleNodeGraph failed. Expected %v, got %v", expected, result)
	}
}

// Test Case 10: No Path Exists
func TestNoPathExists(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(3, 4)

	startNode := 0
	endNode := 4
	expected := [][]int{}

	result := graph.FindAllPaths(startNode, endNode)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNoPathExists failed. Expected %v, got %v", expected, result)
	}
}
