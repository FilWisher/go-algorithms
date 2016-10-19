package graph_test

import (
	"testing"
	"github.com/filwisher/algo/graph"
	"github.com/filwisher/algo/set"
)

func TestGraph(t *testing.T) {

	conn := 1
	g := graph.NewGraph()
	for i := 2; i < 6; i++ {
		g.AddEdge(conn, i)
	}

	s := set.NewSet()
	for e := range g.Adjacent(conn) {
		s.Add(e)	
	}
	
	for i := 2; i < 6; i++ {
		if !s.Contains(i) {
			t.Errorf("graph does not contain edge between %d and %d", conn, i)
		}
		for e := range g.Adjacent(i) {
			if e != conn {
				t.Errorf("edges are not bidirectional")	
			}
		}
	}
}
