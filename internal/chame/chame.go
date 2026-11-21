package chame

import (
	"github.com/fredcamps/chamego/internal/graph"
)

const (
	DefaultAlgorithm = "BFS"
)

type Settings struct {
	SearchAlgorithm string
}

type NotFoundError struct {
	Message string
}

type IntegrityError struct {
	Message string
}

// create façade for package gragh to perform operations insert, update, delete and search on graph nodes
type ChameGraph struct {
	Node     *graph.GraphNode[any]
	Directed bool
	Settings Settings
	nodes    []*graph.GraphNode[any]
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// initialize a new graph
func Chame(directed bool) *ChameGraph {
	chameGraph := &ChameGraph{}
	chameGraph.Directed = directed
	chameGraph.Settings.SearchAlgorithm = DefaultAlgorithm
	return chameGraph
}

// perform set to include key, value, relations in the graph node using directed with defsult value False
func (chameGraph *ChameGraph) Set(key string, value any, relations ...*graph.GraphNode[any]) *graph.GraphNode[any] {
	// check if key exists in graph and raise

	chameGraph.Node = graph.New(key, value, chameGraph.Directed)
	chameGraph.nodes = append(chameGraph.nodes, chameGraph.Node)
	for _, relation := range relations {
		chameGraph.Node.AddRelation(relation)
	}
	return chameGraph.Node
}

// selectSearchAlgorithm chooses the search algorithm based on settings.
func (chameGraph *ChameGraph) selectSearchAlgorithm(node *graph.GraphNode[any], key string) *graph.GraphNode[any] {
	if chameGraph.Settings.SearchAlgorithm == "BFS" {
		return node.FindNodeByKeyBFS(key)
	}
	return node.FindNodeByKeyDFS(key)
}

// perform get to retrieve the graph node by key
func (chameGraph *ChameGraph) Get(key string) (*graph.GraphNode[any], error) {
	var result *graph.GraphNode[any]
	if chameGraph.Node != nil {
		result = chameGraph.selectSearchAlgorithm(chameGraph.Node, key)
		if result != nil {
			return result, nil
		}
	}

	for _, node := range chameGraph.nodes {
		if node == nil {
			continue
		}
		result = chameGraph.selectSearchAlgorithm(node, key)
		if result != nil {
			return result, nil
		}
	}
	return nil, &NotFoundError{Message: "Node not found: " + key}
}

// perform delete to remove the graph node by key
func (chameGraph *ChameGraph) Delete(key string) (*graph.GraphNode[any], error) {
	node, err := chameGraph.Get(key)
	if err != nil {
		return nil, err
	}
	node.Destroy()
	for i, knownNode := range chameGraph.nodes {
		if knownNode == node {
			chameGraph.nodes = append(chameGraph.nodes[:i], chameGraph.nodes[i+1:]...)
			break
		}
	}
	if chameGraph.Node == node {
		chameGraph.Node = nil
		if len(chameGraph.nodes) > 0 {
			chameGraph.Node = chameGraph.nodes[len(chameGraph.nodes)-1]
		}
	}
	return node, nil
}

// perform update on graph node by key and value
func (chameGraph *ChameGraph) Update(key string, value any, relations ...*graph.GraphNode[any]) (*graph.GraphNode[any], error) {
	node, err := chameGraph.Get(key)
	if err != nil {
		return nil, err
	}
	node.Update(key, value, relations...)
	return node, nil
}
