package graph_test

import (
	"github.com/filwisher/algo/graph"
	"github.com/filwisher/algo/set"
	"testing"
)

func TestBFS(t *testing.T) {

	g := graph.NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(1, 4)

	src := 1
	path := graph.NewBFSPath(&g, src)
	
	s := set.NewSet()
	for k := range path.Parent {
		if s.Contains(k) {
			t.Errorf("node %d visited more than once", k)
		}
		s.Add(k)
	}
	
	for i := 1; i < 6; i++ {
		if !s.Contains(i) && i != src {
			t.Errorf("node %d not visited", i)	
		}	
	}
}
