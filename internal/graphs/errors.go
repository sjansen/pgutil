package graphs

import "fmt"

type InvalidEdgeError struct {
	Node string
	Edge string
}

func (e *InvalidEdgeError) Error() string {
	return fmt.Sprintf("node %q references non-existent node %q", e.Node, e.Edge)
}
