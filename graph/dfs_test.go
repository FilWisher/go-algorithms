package graph_test

import (
	"github.com/filwisher/algo/graph"
	"github.com/filwisher/algo/set"
	"testing"
)

func TestDFS(t *testing.T) {

	g := graph.NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(3, 4)
	g.AddEdge(2, 4)

	src := 1
	path := graph.NewDFSPath(&g, src)
	
	s := set.NewSet()
	for k := range path.Parent {
		if s.Contains(k) {
			t.Errorf("node %d visited twice", k)
		}
		s.Add(k)
	}
	for i := 1; i < 5; i++ {
		if !s.Contains(i) && i != src {
			t.Errorf("dfs path does not contain %d", i)	
		}	
	}
}
