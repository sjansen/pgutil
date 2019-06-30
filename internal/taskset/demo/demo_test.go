package demo_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/taskset/demo"
)

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	buffer := &bytes.Buffer{}
	factory := &demo.TargetFactory{
		Stdout: buffer,
	}
	tmp := factory.NewTarget()
	require.NotNil(tmp)

	target := tmp.(*demo.Target)
	target.String = "!abbcF"

	ctx := context.TODO()
	require.NoError(target.Ready())
	require.Equal("", buffer.String())

	t1 := &demo.Echo{}
	require.NoError(t1.Ready())
	require.NoError(
		target.Run(ctx, t1),
	)
	require.Equal("!abbcF\n", buffer.String())

	t2 := &demo.Rev{}
	require.NoError(t2.Ready())
	require.NoError(
		target.Run(ctx, t2),
	)
	require.Equal("Fcbba!", target.String)

	t3 := &demo.Rot13{}
	require.NoError(t3.Ready())
	require.NoError(
		target.Run(ctx, t3),
	)
	require.Equal("Spoon!", target.String)

	t4 := &demo.Echo{}
	require.NoError(t4.Ready())
	require.NoError(
		target.Run(ctx, t4),
	)
	require.Equal(
		"!abbcF\nSpoon!\n",
		buffer.String(),
	)
}
