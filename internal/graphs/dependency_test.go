package graphs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/graphs"
)

var acyclic = map[string][]string{
	"a": {"b", "m", "u"},
	"b": {"o", "u"},
	"e": {"b", "l"},
	"g": {"k", "l", "m", "n", "o"},
	"i": {"k", "a", "r", "t"},
	"k": {"a", "b"},
	"l": {"b", "t"},
	"m": {"e", "r"},
	"n": {"a", "e", "i", "o", "u"},
	"o": {"r"},
	"r": {"t"},
	"t": {},
	"u": {"o", "r"},
}

var cyclic = map[string][]string{
	"foo": {"bar", "baz"},
	"bar": {},
	"baz": {},
	"0":   {"1"},
	"1":   {"2"},
	"2":   {"3"},
	"3":   {"4"},
	"4":   {"5"},
	"5":   {"6"},
	"6":   {"7"},
	"7":   {"8"},
	"8":   {"9"},
	"9":   {"0"},
}

var reversed = map[string][]string{
	"a": {"k", "i", "n", "g"},
	"b": {"k", "l", "m", "n"},
	"e": {"m", "n"},
	"g": {},
	"i": {"n"},
	"k": {"i", "n"},
	"l": {"a", "e"},
	"m": {"a", "k", "g"},
	"n": {"g"},
	"o": {"k", "u"},
	"r": {"a", "e", "i", "o", "u"},
	"t": {"r"},
	"u": {"a", "b"},
}

var unpredictable = map[string][]string{
	"a": {},
	"b": {},
	"c": {"h", "j"},
	"d": {},
	"e": {"k"},
	"f": {},
	"g": {"k", "h"},
	"h": {"i"},
	"i": {},
	"j": {},
	"k": {"b", "i", "j"},
	"l": {},
	"m": {},
}

func TestNewDependencyGraph(t *testing.T) {
	require := require.New(t)

	expected := &graphs.InvalidEdgeError{Node: "baz", Edge: "qux"}
	nodes := map[string][]string{
		"foo": {"bar", "baz"},
		"bar": {"baz"},
		"baz": {"qux"},
	}
	g, err := graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(nodes),
	)
	require.Nil(g)
	require.Equal(expected, err)
	require.NotEmpty(err.Error())
}

func TestTSort(t *testing.T) {
	require := require.New(t)

	expected := []string{"t", "r", "o", "u", "b", "l", "e", "m", "a", "k", "i", "n", "g"}
	g, err := graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(acyclic),
	)
	require.NoError(err)
	actual, cycle := g.TSort()
	require.Equal(expected, actual)
	require.Empty(cycle)

	expected = []string{"g", "n", "i", "k", "a", "m", "e", "l", "b", "u", "o", "r", "t"}
	g, err = graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(reversed),
	)
	require.NoError(err)
	actual, cycle = g.TSort()
	require.Equal(expected, actual)
	require.Empty(cycle)

	expected = []string{"a", "b", "i", "h", "j", "c", "d", "k", "e", "f", "g", "l", "m"}
	g, err = graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(unpredictable),
	)
	require.NoError(err)
	actual, cycle = g.TSort()
	require.Equal(expected, actual)
	require.Empty(cycle)

	g, err = graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(cyclic),
	)
	require.NoError(err)
	actual, cycle = g.TSort()
	require.Empty(actual)
	for _, n := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		require.Contains(cycle, n)
	}
}
