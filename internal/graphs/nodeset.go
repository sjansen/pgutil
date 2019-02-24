package graphs

type NodeSet map[string]struct{}

func (nodes NodeSet) Add(node string) {
	nodes[node] = struct{}{}
}

func (nodes NodeSet) Remove(node string) {
	delete(nodes, node)
}

func (nodes NodeSet) Contains(node string) bool {
	_, ok := nodes[node]
	return ok
}

func (nodes NodeSet) Size() int {
	return len(nodes)
}
