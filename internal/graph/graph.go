package graph

import (
	"slices"
	"sync"
)

// a thread-safe in-memory graph struct with directed and undirected relations with key as string attribute and value as generic type
type GraphNode[T any] struct {
	Directed  bool
	Key       string
	Value     T
	Relations []*GraphNode[T]
	mu        sync.RWMutex
}

// initialize a new graph node with the given key, value, and directed flag
func New[T any](key string, value T, directed bool) *GraphNode[T] {
	node := &GraphNode[T]{
		Key:      key,
		Value:    value,
		Directed: directed,
	}
	return node
}

// add a bidirectional relation considering the directed flag of the current node
func (node *GraphNode[T]) AddRelation(relation *GraphNode[T]) {
	if relation == nil {
		return
	}

	node.mu.Lock()
	defer node.mu.Unlock()
	for _, existing := range node.Relations {
		if existing == relation {
			return
		}
	}
	node.Relations = append(node.Relations, relation)

	if !node.Directed {
		relation.mu.Lock()
		defer relation.mu.Unlock()
		if slices.Contains(relation.Relations, node) {
			return
		}
		relation.Relations = append(relation.Relations, node)
	}
}

// remove a bidirectional relation considering Directed flag
func (node *GraphNode[T]) RemoveRelation(relation *GraphNode[T]) {
	if relation == nil {
		return
	}

	node.mu.Lock()
	defer node.mu.Unlock()
	for i, existing := range node.Relations {
		if existing == relation {
			node.Relations = slices.Delete(node.Relations, i, i+1)
			break
		}
	}

	if !node.Directed {
		relation.mu.Lock()
		defer relation.mu.Unlock()
		for i, existing := range relation.Relations {
			if existing == node {
				relation.Relations = slices.Delete(relation.Relations, i, i+1)
				break
			}
		}
	}
}

// perform a breadth-first search traversal of the graph starting from the current node finding node by key
func (node *GraphNode[T]) FindNodeByKeyBFS(key string) *GraphNode[T] {
	visited := make(map[string]bool)
	return bfsFind(node, key, visited)
}

// perform bfsFind recursively to find the node by key
func bfsFind[T any](node *GraphNode[T], key string, visited map[string]bool) *GraphNode[T] {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	}
	visited[node.Key] = true

	node.mu.RLock()
	relations := append([]*GraphNode[T](nil), node.Relations...)
	node.mu.RUnlock()

	for _, relation := range relations {
		if !visited[relation.Key] {
			found := bfsFind(relation, key, visited)
			if found != nil {
				return found
			}
		}
	}
	return nil
}

// perform depth-first search traversal of the graph starting from the current node finding node by key
func (node *GraphNode[T]) FindNodeByKeyDFS(key string) *GraphNode[T] {
	visited := make(map[string]bool)
	return dfsFind(node, key, visited)
}

// perform dfsFind recursively to find the node by key
func dfsFind[T any](node *GraphNode[T], key string, visited map[string]bool) *GraphNode[T] {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	}
	visited[node.Key] = true

	node.mu.RLock()
	relations := append([]*GraphNode[T](nil), node.Relations...)
	node.mu.RUnlock()

	for _, relation := range relations {
		if !visited[relation.Key] {
			found := dfsFind(relation, key, visited)
			if found != nil {
				return found
			}
		}
	}
	return nil
}

// destroy the graph node and its relations considering the directed flag of the current node
func (node *GraphNode[T]) Destroy() {
	node.mu.RLock()
	relations := append([]*GraphNode[T](nil), node.Relations...)
	directed := node.Directed
	node.mu.RUnlock()

	if !directed {
		for _, relation := range relations {
			relation.RemoveRelation(node)
		}
	}

	node.mu.Lock()
	node.Relations = slices.Delete(node.Relations, 0, len(node.Relations))
	node.mu.Unlock()
}

// Update graph properties with new key, value, relations, maintaining thread safety
func (node *GraphNode[T]) Update(key string, value T, relations ...*GraphNode[T]) {
	node.mu.RLock()
	oldRelations := append([]*GraphNode[T](nil), node.Relations...)
	node.mu.RUnlock()

	for _, relation := range oldRelations {
		node.RemoveRelation(relation)
	}

	node.mu.Lock()
	node.Key = key
	node.Value = value
	node.mu.Unlock()

	for _, relation := range relations {
		node.AddRelation(relation)
	}
}
