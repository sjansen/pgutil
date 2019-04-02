package dag_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/dag"
)

func TestNodeSet(t *testing.T) {
	require := require.New(t)

	var set dag.NodeSet
	require.Equal(0, set.Size())
	require.False(set.Contains("foo"))
	set.Remove("bar")

	set = dag.NodeSet{}
	require.Equal(0, set.Size())
	set.Add("foo")
	set.Add("bar")
	require.Equal(2, set.Size())
	require.True(set.Contains("foo"))
	require.True(set.Contains("bar"))
	set.Remove("bar")
	require.Equal(1, set.Size())
	require.True(set.Contains("foo"))
	require.False(set.Contains("bar"))
}
