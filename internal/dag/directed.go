package dag

type DirectedGraph map[string]NodeSet

// NewDirectedGraph converts a list of nodes and edges into a graph
func NewDirectedGraph(nodes map[string][]string) (*DirectedGraph, error) {
	g := DirectedGraph{}
	for n := range nodes {
		set := NodeSet{}
		for _, m := range nodes[n] {
			if _, ok := nodes[m]; !ok {
				err := &InvalidEdgeError{Node: n, Edge: m}
				return nil, err
			}
			set.Add(m)
		}
		g[n] = set
	}
	return &g, nil
}

// AddEdge adds a directed relationship from one node to another to the graph
func (g *DirectedGraph) AddEdge(node, edge string) {
	nodes := *g
	if set, ok := nodes[node]; ok {
		set.Add(edge)
	} else {
		nodes[node] = NodeSet{edge: struct{}{}}
	}
	if _, ok := nodes[edge]; !ok {
		nodes[edge] = NodeSet{}
	}
}

// AddEdges adds a directed relationship from one node to multiple others in the graph
func (g *DirectedGraph) AddEdges(node string, edges []string) {
	for _, edge := range edges {
		g.AddEdge(node, edge)
	}
}

// HasNode checks if a node is included in the graph
func (g *DirectedGraph) HasNode(node string) bool {
	nodes := *g
	_, ok := nodes[node]
	return ok
}

// RemoveEdge removes a relationship from one node to another from the graph
func (g *DirectedGraph) RemoveEdge(node, edge string) {
	nodes := *g
	delete(nodes[node], edge)
}

// RemoveNode remove a node and its outbound edges from the graph
func (g *DirectedGraph) RemoveNode(node string) {
	nodes := *g
	for n := range nodes {
		g.RemoveEdge(n, node)
	}
	delete(nodes, node)
}

// OutDegree counts the number of edges outbound from a node in the graph
func (g *DirectedGraph) OutDegree(node string) int {
	return len((*g)[node])
}
