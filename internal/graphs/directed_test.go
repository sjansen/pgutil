package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/graphs"
)

func TestNewDirectedGraph(t *testing.T) {
	require := require.New(t)

	expected := &graphs.InvalidEdgeError{Node: "baz", Edge: "qux"}
	nodes := map[string][]string{
		"foo": {"bar", "baz"},
		"bar": {"baz"},
		"baz": {"qux"},
	}
	g, err := graphs.NewDirectedGraph(nodes)
	require.Nil(g)
	require.Equal(expected, err)
	require.NotEmpty(err.Error())
}

func TestDirectedGraph(t *testing.T) {
	require := require.New(t)

	// nil constructor
	g, err := graphs.NewDirectedGraph(nil)
	require.NoError(err)
	g.AddEdges("foo", []string{"bar", "baz"})
	require.True(g.HasNode("foo"))
	require.True(g.HasNode("bar"))
	require.True(g.HasNode("bar"))
	require.Equal(2, g.OutDegree("foo"))
	require.Equal(0, g.OutDegree("bar"))
	require.Equal(0, g.OutDegree("baz"))

	// non-nil constructor
	nodes := map[string][]string{
		"foo": {"baz"},
		"bar": {},
		"baz": {"foo", "bar"},
	}
	g, err = graphs.NewDirectedGraph(nodes)
	require.NoError(err)
	require.Equal(1, g.OutDegree("foo"))
	require.Equal(0, g.OutDegree("bar"))
	require.Equal(2, g.OutDegree("baz"))

	// increase number of edges
	g.AddEdge("foo", "bar")
	require.Equal(2, g.OutDegree("foo"))

	// remove all edges
	g.RemoveEdge("bar", "baz")
	require.Equal(0, g.OutDegree("bar"))
	require.True(g.HasNode("bar"))

	// add new nodes
	g.AddEdge("qux", "quux")
	require.Equal(1, g.OutDegree("qux"))
	require.Equal(0, g.OutDegree("quux"))
	require.True(g.HasNode("quux"))

	// remove nodes
	g.RemoveNode("baz")
	require.Equal(1, g.OutDegree("foo"))
	require.False(g.HasNode("baz"))
	g.RemoveNode("quux")
	require.Equal(0, g.OutDegree("qux"))
	require.Equal(0, g.OutDegree("quux"))
	require.False(g.HasNode("quux"))

	// remove non-existent nodes
	g.RemoveEdge("corge", "grault")
	require.Equal(0, g.OutDegree("corge"))
	require.Equal(0, g.OutDegree("grault"))
	require.False(g.HasNode("corge"))
	require.False(g.HasNode("grault"))
	g.RemoveNode("garply")
	require.Equal(0, g.OutDegree("garply"))
	require.False(g.HasNode("garply"))
}
