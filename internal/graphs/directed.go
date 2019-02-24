package graphs

type DirectedGraph map[string]NodeSet

type NodeSet map[string]struct{}

func NewDirectedGraph(nodes map[string][]string) (*DirectedGraph, error) {
	g := DirectedGraph{}
	for n := range nodes {
		h := NodeSet{}
		for _, m := range nodes[n] {
			if _, ok := nodes[m]; !ok {
				err := &InvalidEdgeError{Node: n, Edge: m}
				return nil, err
			}
			h[m] = struct{}{}
		}
		g[n] = h
	}
	return &g, nil
}

func (g *DirectedGraph) AddEdge(node, edge string) {
	nodes := *g
	if _, ok := nodes[node]; !ok {
		nodes[node] = NodeSet{edge: struct{}{}}
	} else {
		nodes[node][edge] = struct{}{}
	}
	if _, ok := nodes[edge]; !ok {
		nodes[edge] = NodeSet{}
	}
}

func (g *DirectedGraph) AddEdges(node string, edges []string) {
	for _, edge := range edges {
		g.AddEdge(node, edge)
	}
}

func (g *DirectedGraph) HasNode(node string) bool {
	nodes := *g
	_, ok := nodes[node]
	return ok
}

func (g *DirectedGraph) RemoveEdge(node, edge string) {
	nodes := *g
	delete(nodes[node], edge)
}

func (g *DirectedGraph) RemoveNode(node string) {
	nodes := *g
	for n := range nodes {
		g.RemoveEdge(n, node)
	}
	delete(nodes, node)
}

func (g *DirectedGraph) OutDegree(node string) int {
	return len((*g)[node])
}
