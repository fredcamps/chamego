package tests

import (
	"slices"
	"testing"

	"github.com/fredcamps/chamego/internal/chame"
)

// test ChameGraph creation and basic operations as an undirected graph
func TestChameGraph(t *testing.T) {
	// Create a new graph node
	graph := chame.Chame(false)

	// Set new values for the graph node
	nodeA := graph.Set("A1", "Node A1")
	if graph.Node.Key != "A1" || graph.Node.Value != "Node A1" {
		t.Errorf("Expected key 'A1' and value 'Node A1', got key '%s' and value '%v'", graph.Node.Key, graph.Node.Value)
	}

	// Set a relation to another graph node
	nodeB := graph.Set("B", "Node B", nodeA)
	nodeC := graph.Set("C", "Node C", nodeA, nodeB)

	// perform nodeA destroy test
	graph.Delete("A1")
	foundA, _ := graph.Get(nodeA.Key)
	if foundA != nil || slices.Contains(nodeC.Relations, nodeA) {
		t.Errorf("Expected node 'A1' to be destroyed, but it still exists")
	}

	// Find a node by key
	foundNode, _ := graph.Get("B")
	if foundNode == nil || foundNode.Key != "B" {
		t.Errorf("Expected to find node 'B', got: %v", foundNode)
	}

	// Delete a relation
	graph.Delete("B")
	if len(graph.Node.Relations) != 0 {
		t.Errorf("Expected no relations after deletion, got: %v", graph.Node.Relations)
	}

	// Check if nodeC has no relations
	graph.Get("C")
	if graph.Node == nodeC && len(nodeC.Relations) != 0 {
		t.Errorf("Expected nodeC to have no relations, got: %v", nodeC.Relations)
	}

	// Create new node and then do nodeC Update test with new Relation as argument
	nodeX := graph.Set("NodeX", "Last Node")
	if _, err := graph.Update("C", "Another C-level value", nodeX); err != nil {
		t.Fatalf("Expected to update node 'C', got error: %v", err)
	}

	// Assert that nodeC has nodeX as a relation after the update
	if !slices.Contains(nodeC.Relations, nodeX) {
		t.Errorf("Expected nodeC to have nodeX as a relation after update, got: %v", nodeC.Relations)
	}

	// Validate trying to delete a not present node key
	_, err := graph.Delete("NonExistentKey")
	if err == nil {
		t.Errorf("Expected error when deleting non-existent node key, got nil")
	}
}

// test ChameGraph creation and basic operations as a directed graph scenario.
func TestChameGraphDirected(t *testing.T) {
	// Create a new directed graph node
	graph := chame.Chame(true)

	// Set new values for the graph node
	nodeA := graph.Set("A1", "Node A1")
	if graph.Node.Key != "A1" || graph.Node.Value != "Node A1" {
		t.Errorf("Expected key 'A1' and value 'Node A1', got key '%s' and value '%v'", graph.Node.Key, graph.Node.Value)
	}

	// Set a relation to another graph node
	nodeB := graph.Set("B", "Node B", nodeA)
	nodeC := graph.Set("C", "Node C", nodeA, nodeB)

	// perform nodeA destroy test
	graph.Delete("A1")
	foundA, _ := graph.Get(nodeA.Key)
	if foundA != nil {
		t.Errorf("Expected node 'A1' to be destroyed, but it still exists")
	}

	// Find a node by key
	foundNode, _ := graph.Get("B")
	if foundNode == nil || foundNode.Key != "B" {
		t.Errorf("Expected to find node 'B', got: %v", foundNode)
	}

	// Delete a relation
	graph.Delete("B")
	if len(graph.Node.Relations) != 0 {
		t.Errorf("Expected no relations after deletion, got: %v", graph.Node.Relations)
	}

	// Check if nodeC has no relations
	graph.Get("C")
	if graph.Node == nodeC && len(nodeC.Relations) != 0 {
		t.Errorf("Expected nodeC to have no relations, got: %v", nodeC.Relations)
	}

	// Create new node and then do nodeC Update test with new Relation as argument
	nodeX := graph.Set("NodeX", "Last Node")
	if _, err := graph.Update("C", "Another C-level value", nodeX); err != nil {
		t.Fatalf("Expected to update node 'C', got error: %v", err)
	}

	// Assert that nodeC has nodeX as a relation after the update
	if !slices.Contains(nodeC.Relations, nodeX) {
		t.Errorf("Expected nodeC to have nodeX as a relation after update, got: %v", nodeC.Relations)
	}

	// Validate trying to delete a not present node key
	_, err := graph.Delete("NonExistentKey")
	if err == nil {
		t.Errorf("Expected error when deleting non-existent node key, got nil")
	}
}
