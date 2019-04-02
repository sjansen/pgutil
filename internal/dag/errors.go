package dag

import "fmt"

type CircularDependencyError struct {
	Cycle []string
}

func (e *CircularDependencyError) Error() string {
	return fmt.Sprintf("circular dependency detected: %q", e.Cycle)
}

type InvalidEdgeError struct {
	Node string
	Edge string
}

func (e *InvalidEdgeError) Error() string {
	return fmt.Sprintf("node %q references non-existent node %q", e.Node, e.Edge)
}
