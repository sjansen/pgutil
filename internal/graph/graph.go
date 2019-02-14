package graph

type DirectedGraph struct {
	nodes map[string][]string
}

func NewDirectedGraph(g map[string][]string) (*DirectedGraph, error) {
	for n := range g {
		for _, m := range g[n] {
			if _, ok := g[m]; !ok {
				err := &InvalidEdgeError{Node: n, Edge: m}
				return nil, err
			}
		}
	}
	return &DirectedGraph{nodes: g}, nil
}

func (g *DirectedGraph) TSort() (sorted, cycle []string) {
	v := &visitor{
		graph:    g.nodes,
		sorted:   make([]string, 0, len(g.nodes)),
		visited:  map[string]struct{}{},
		visiting: map[string]struct{}{},
	}
	for n := range g.nodes {
		if _, ok := v.visited[n]; ok {
			continue
		}
		v.visit(n)
		if v.cycle != nil {
			return nil, v.cycle
		}
	}
	return v.sorted, nil
}

type visitor struct {
	graph    map[string][]string
	visited  map[string]struct{}
	visiting map[string]struct{}
	cycle    []string
	sorted   []string
}

func (v *visitor) visit(n string) (cycleStart string) {
	if _, ok := v.visiting[n]; ok {
		if v.cycle == nil {
			v.cycle = make([]string, 0)
		}
		return n
	} else if _, ok := v.visited[n]; ok {
		return ""
	}
	v.visiting[n] = struct{}{}
	for _, m := range v.graph[n] {
		cycleStart := v.visit(m)
		if v.cycle != nil {
			if cycleStart != "" {
				if n == cycleStart {
					cycleStart = ""
				}
				v.cycle = append(v.cycle, n)
			}
			return cycleStart
		}
	}
	delete(v.visiting, n)
	v.visited[n] = struct{}{}
	v.sorted = append(v.sorted, n)
	return ""
}
