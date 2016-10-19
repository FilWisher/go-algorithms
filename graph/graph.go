package graph

import (
	"github.com/filwisher/algo/set"
)

type Graph struct {
	Edges map[int] *set.Set
}

func NewGraph() Graph {
	return Graph{
		Edges: make(map[int] *set.Set),
	}
}

func (g *Graph) addSingleEdge(u, v int) {
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = set.NewSet()
	}
	g.Edges[u].Add(v)
}

func (g *Graph) AddEdge(u, v int) {
	g.addSingleEdge(u, v)
	g.addSingleEdge(v, u)
}

func (g *Graph) Adjacent(u int) <-chan int {
	ch := make(chan int)
	go func() {
		if set, ok := g.Edges[u]; ok {
			for e := range set.Each() {
				ch <- e.(int)
			}	
			close(ch)
		}
	}()
	return ch
}
