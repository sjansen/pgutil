package dag

type DependencyGraph struct {
	dependencies map[string]NodeSet
	dependents   map[string]NodeSet
	pending      *pendingTask
}

type pendingTask struct {
	id   string
	next *pendingTask
}

func NewDependencyGraph(nodes map[string][]string) (*DependencyGraph, error) {
	digraph, err := NewDirectedGraph(nodes)
	if err != nil {
		return nil, err
	}

	g := &DependencyGraph{
		dependencies: map[string]NodeSet{},
		dependents:   map[string]NodeSet{},
	}
	nodeOrder, cycle := TSort(digraph)
	if cycle != nil {
		err := &CircularDependencyError{
			Cycle: cycle,
		}
		return nil, err
	}
	var pending *pendingTask
	for i := len(nodeOrder) - 1; i >= 0; i-- {
		pending = &pendingTask{
			id:   nodeOrder[i],
			next: pending,
		}
	}
	g.pending = pending

	for n, edges := range *digraph {
		g.dependencies[n] = NodeSet{}
		for m := range edges {
			g.dependencies[n].Add(m)
			if g.dependents[m] == nil {
				g.dependents[m] = NodeSet{}
			}
			g.dependents[m].Add(n)
		}
	}

	return g, nil
}

func (g *DependencyGraph) HasPending() bool {
	return g.pending != nil
}

func (g *DependencyGraph) Next(completed string) []string {
	var ready []string
	for id := range g.dependents[completed] {
		g.dependencies[id].Remove(completed)
	}

	head := g.pending
	for head != nil && g.dependencies[head.id].Size() < 1 {
		ready = append(ready, head.id)
		head = head.next
	}
	g.pending = head

	for head = g.pending; head != nil && head.next != nil; head = head.next {
		p := head.next
		if g.dependencies[p.id].Size() < 1 {
			ready = append(ready, p.id)
			head.next = p.next
		}
	}
	return ready
}
