package graphs

type DirectedGraph map[string]NodeSet

type NodeSet map[string]struct{}

func NewDirectedGraph(nodes map[string][]string) *DirectedGraph {
	g := DirectedGraph{}
	for n := range nodes {
		h := NodeSet{}
		for _, m := range nodes[n] {
			h[m] = struct{}{}
		}
		g[n] = h
	}
	return &g
}
