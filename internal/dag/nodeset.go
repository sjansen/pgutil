package dag

// A NodeSet is a collection of nodes
type NodeSet map[string]struct{}

// Add adds a node to the set
func (nodes NodeSet) Add(node string) {
	nodes[node] = struct{}{}
}

// Remove removes a node from the set
func (nodes NodeSet) Remove(node string) {
	delete(nodes, node)
}

// Contains checks if a specific node is in the set
func (nodes NodeSet) Contains(node string) bool {
	_, ok := nodes[node]
	return ok
}

// Size counts the number of nodes in the set
func (nodes NodeSet) Size() int {
	return len(nodes)
}
