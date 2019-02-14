package graph

import "sort"

type DependencyGraph struct {
	nodes  []string
	edges  map[string][]string
	sorted bool
}

func NewDependencyGraph(g map[string][]string) (*DependencyGraph, error) {
	nodes := make([]string, 0, len(g))
	edges := make(map[string][]string, len(g))
	for n := range g {
		nodes = append(nodes, n)
		edges[n] = make([]string, 0, len(g[n]))
		for _, m := range g[n] {
			edges[n] = append(edges[n], m)
			if _, ok := g[m]; !ok {
				err := &InvalidEdgeError{Node: n, Edge: m}
				return nil, err
			}
		}
	}
	return &DependencyGraph{nodes: nodes, edges: edges}, nil
}

func (g *DependencyGraph) TSort() (nodes, cycle []string) {
	if !g.sorted {
		sort.Strings(g.nodes)
		for _, n := range g.nodes {
			sort.Strings(g.edges[n])
		}
		g.sorted = true
	}

	v := &visitor{
		graph:    g.edges,
		sorted:   make([]string, 0, len(g.nodes)),
		visited:  map[string]struct{}{},
		visiting: map[string]struct{}{},
	}
	for _, n := range g.nodes {
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
