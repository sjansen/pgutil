package strbuf_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/strbuf"
)

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	buffer := &bytes.Buffer{}
	factory := &strbuf.TargetFactory{
		Stdout: buffer,
	}
	target := factory.NewTarget()

	ctx := context.TODO()
	target.(*strbuf.Target).Data = "!abbcF"
	require.NoError(target.Analyze())
	require.Equal("", buffer.String())

	t1 := &strbuf.Echo{}
	require.NoError(t1.Check())
	require.NoError(
		target.Handle(ctx, t1),
	)
	require.Equal("!abbcF\n", buffer.String())

	t2 := &strbuf.Rev{}
	require.NoError(t2.Check())
	require.NoError(
		target.Handle(ctx, t2),
	)
	require.Equal("Fcbba!", target.(*strbuf.Target).Data)

	t3 := &strbuf.Rot13{}
	require.NoError(t3.Check())
	require.NoError(
		target.Handle(ctx, t3),
	)
	require.Equal("Spoon!", target.(*strbuf.Target).Data)

	t4 := &strbuf.Echo{}
	require.NoError(t4.Check())
	require.NoError(
		target.Handle(ctx, t4),
	)
	require.Equal(
		"!abbcF\nSpoon!\n",
		buffer.String(),
	)
}
