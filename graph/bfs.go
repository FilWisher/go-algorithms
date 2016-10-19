package graph

import (
	"github.com/filwisher/algo/queue"
)

type BFSPath struct {
	Src int
	G *Graph
	Dist map[int] int
	Parent map[int] int
}

func NewBFSPath(g *Graph, v int) BFSPath {
	path := BFSPath{
		Src: v,
		G: g,
		Dist: make(map[int] int),
		Parent: make(map[int] int),
	}
	path.bfs()
	return path
}

func (p *BFSPath) bfs() {
	q := queue.NewQueue()
	p.Dist[p.Src] = 0
	q.Enqueue(p.Src)
	for !q.IsEmpty() {
		curr, _ := q.Dequeue()
		for u := range p.G.Adjacent(curr.(int)) {
			if _, ok := p.Dist[u]; !ok {
				p.Dist[u] = p.Dist[curr.(int)] + 1
				p.Parent[u] = curr.(int)
				q.Enqueue(u)
			}	
		}
	}
}
