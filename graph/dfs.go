package graph


type DFSPath struct {
	Src int
	Visited map[int] bool
	Parent map[int] int
	G *Graph
}

func NewDFSPath(g *Graph, v int) *DFSPath {
	path := &DFSPath{
		Src: v,
		Visited: make(map[int] bool),
		Parent: make(map[int] int),
		G: g,
	}
	path.dfs(v)
	return path
}

func (p *DFSPath) dfs(v int) {
	p.Visited[v] = true
	for u := range p.G.Adjacent(v) {
		if !p.Visited[u] {
			p.Parent[u] = v
			p.dfs(u)	
		}
	}
}
