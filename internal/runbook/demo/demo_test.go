package demo_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/demo"
)

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	buffer := &bytes.Buffer{}
	factory := &demo.TargetFactory{
		Stdout: buffer,
	}
	target := factory.NewTarget()

	ctx := context.TODO()
	target.(*demo.Target).Data = "!abbcF"
	require.NoError(target.Analyze())
	require.Equal("", buffer.String())

	t1 := &demo.Echo{}
	require.NoError(t1.Check())
	require.NoError(
		target.Handle(ctx, t1),
	)
	require.Equal("!abbcF\n", buffer.String())

	t2 := &demo.Rev{}
	require.NoError(t2.Check())
	require.NoError(
		target.Handle(ctx, t2),
	)
	require.Equal("Fcbba!", target.(*demo.Target).Data)

	t3 := &demo.Rot13{}
	require.NoError(t3.Check())
	require.NoError(
		target.Handle(ctx, t3),
	)
	require.Equal("Spoon!", target.(*demo.Target).Data)

	t4 := &demo.Echo{}
	require.NoError(t4.Check())
	require.NoError(
		target.Handle(ctx, t4),
	)
	require.Equal(
		"!abbcF\nSpoon!\n",
		buffer.String(),
	)
}
