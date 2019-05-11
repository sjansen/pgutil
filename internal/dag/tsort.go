package dag

import "sort"

// TSort produces a topological sort of a directed graph
func TSort(g *DirectedGraph) (result, cycle []string) {
	digraph := *g
	nodes := make([]string, 0, len(digraph))
	edges := make(map[string][]string, len(digraph))

	for n, nodeSet := range digraph {
		nodes = append(nodes, n)
		tmp := make([]string, 0, len(nodeSet))
		for m := range nodeSet {
			tmp = append(tmp, m)
		}
		sort.Strings(tmp)
		edges[n] = tmp
	}
	sort.Strings(nodes)

	v := &visitor{
		graph:    edges,
		sorted:   make([]string, 0, len(nodes)),
		visited:  map[string]struct{}{},
		visiting: map[string]struct{}{},
	}
	for _, n := range nodes {
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
