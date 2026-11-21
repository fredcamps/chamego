package tests

import (
	"testing"

	"github.com/fredcamps/chamego/internal/graph"
)

func newTestGraphNode[T any](key string, value T, directed bool) *graph.GraphNode[T] {
	return graph.New(key, value, directed)
}

func TestGraphNodeNew(t *testing.T) {
	node := newTestGraphNode("A", "value", true)

	if node.Key != "A" || node.Value != "value" || !node.Directed {
		t.Fatalf("New() = {Key: %q, Value: %q, Directed: %t}, want {Key: %q, Value: %q, Directed: true}",
			node.Key, node.Value, node.Directed, "A", "value")
	}
	if len(node.Relations) != 0 {
		t.Fatalf("New() created %d relations, want none", len(node.Relations))
	}
}

func TestGraphNodeAddAndRemoveUndirectedRelation(t *testing.T) {
	a := newTestGraphNode("A", "a", false)
	b := newTestGraphNode("B", "b", false)

	a.AddRelation(b)
	a.AddRelation(b)
	a.AddRelation(nil)

	if len(a.Relations) != 1 || a.Relations[0] != b {
		t.Fatalf("A relations = %v, want exactly B", a.Relations)
	}
	if len(b.Relations) != 1 || b.Relations[0] != a {
		t.Fatalf("B relations = %v, want exactly A", b.Relations)
	}

	a.RemoveRelation(b)
	a.RemoveRelation(b)
	a.RemoveRelation(nil)

	if len(a.Relations) != 0 || len(b.Relations) != 0 {
		t.Fatalf("relations after removal: A=%v, B=%v; want none", a.Relations, b.Relations)
	}
}

func TestGraphNodeDirectedRelation(t *testing.T) {
	a := newTestGraphNode("A", "a", true)
	b := newTestGraphNode("B", "b", false)

	a.AddRelation(b)
	if len(a.Relations) != 1 || a.Relations[0] != b {
		t.Fatalf("A relations = %v, want B", a.Relations)
	}
	if len(b.Relations) != 0 {
		t.Fatalf("directed relation unexpectedly added reverse edge: B relations = %v", b.Relations)
	}

	a.RemoveRelation(b)
	if len(a.Relations) != 0 || len(b.Relations) != 0 {
		t.Fatalf("relations after removal: A=%v, B=%v; want none", a.Relations, b.Relations)
	}
}

func TestGraphNodeFindByKey(t *testing.T) {
	a := newTestGraphNode("A", "a", true)
	b := newTestGraphNode("B", "b", true)
	c := newTestGraphNode("C", "c", true)

	a.AddRelation(b)
	b.AddRelation(c)
	c.AddRelation(a)

	for _, find := range []struct {
		name string
		fn   func(string) *graph.GraphNode[string]
	}{
		{name: "BFS", fn: a.FindNodeByKeyBFS},
		{name: "DFS", fn: a.FindNodeByKeyDFS},
	} {
		t.Run(find.name, func(t *testing.T) {
			if got := find.fn("B"); got != b {
				t.Errorf("find B = %v, want node B", got)
			}
			if got := find.fn("C"); got != c {
				t.Errorf("find C = %v, want node C", got)
			}
			if got := find.fn("missing"); got != nil {
				t.Errorf("find missing = %v, want nil", got)
			}
		})
	}
}

func TestGraphNodeFindByKeyNotFound(t *testing.T) {
	a := newTestGraphNode("A", "a", true)
	b := newTestGraphNode("B", "b", true)
	c := newTestGraphNode("C", "c", true)
	d := newTestGraphNode("D", "d", true)
	a.AddRelation(b)
	b.AddRelation(c)

	for _, find := range []struct {
		name string
		fn   func(string) *graph.GraphNode[string]
	}{
		{name: "BFS", fn: a.FindNodeByKeyBFS},
		{name: "DFS", fn: a.FindNodeByKeyDFS},
	} {
		t.Run(find.name, func(t *testing.T) {
			for _, key := range []string{"missing", d.Key} {
				if got := find.fn(key); got != nil {
					t.Errorf("find %q = %v, want nil", key, got)
				}
			}
		})
	}
}

func TestGraphNodeDestroyUndirectedRelations(t *testing.T) {
	a := newTestGraphNode("A", "a", false)
	b := newTestGraphNode("B", "b", false)
	c := newTestGraphNode("C", "c", false)

	a.AddRelation(b)
	a.AddRelation(c)
	a.Destroy()

	if len(a.Relations) != 0 {
		t.Fatalf("A relations after Destroy() = %v, want none", a.Relations)
	}
	if len(b.Relations) != 0 || len(c.Relations) != 0 {
		t.Fatalf("neighbor relations after Destroy(): B=%v, C=%v; want none", b.Relations, c.Relations)
	}
}

func TestGraphNodeDestroyDirectedRelations(t *testing.T) {
	a := newTestGraphNode("A", "a", true)
	b := newTestGraphNode("B", "b", false)

	a.AddRelation(b)
	a.Destroy()

	if len(a.Relations) != 0 {
		t.Fatalf("A relations after Destroy() = %v, want none", a.Relations)
	}
	if len(b.Relations) != 0 {
		t.Fatalf("B relations after destroying directed source = %v, want none", b.Relations)
	}
}

func TestGraphNodeUpdate(t *testing.T) {
	node := newTestGraphNode("A", "old", false)
	oldRelation := newTestGraphNode("old relation", "old", false)
	newRelation := newTestGraphNode("new relation", "new", false)
	node.AddRelation(oldRelation)

	node.Update("B", "new", newRelation)

	if node.Key != "B" || node.Value != "new" {
		t.Fatalf("after Update(), got {Key: %q, Value: %q}, want {Key: %q, Value: %q}",
			node.Key, node.Value, "B", "new")
	}
	if len(node.Relations) != 1 || node.Relations[0] != newRelation {
		t.Fatalf("relations after Update() = %v, want only new relation", node.Relations)
	}
	if len(oldRelation.Relations) != 0 {
		t.Errorf("old relation still points to updated node: %v", oldRelation.Relations)
	}
	if len(newRelation.Relations) != 1 || newRelation.Relations[0] != node {
		t.Errorf("new relation does not point back to updated node: %v", newRelation.Relations)
	}
}
