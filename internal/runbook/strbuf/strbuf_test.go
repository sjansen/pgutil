package strbuf_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/strbuf"
)

func TestRevTask(t *testing.T) {
	require := require.New(t)

	task := &strbuf.RevTask{}
	require.NoError(task.Check())

	actual := task.Munge("!noopS")
	require.Equal("Spoon!", actual)
}

func TestRot13Task(t *testing.T) {
	require := require.New(t)

	task := &strbuf.Rot13Task{}
	require.NoError(task.Check())

	actual := task.Munge("Fcbba!")
	require.Equal("Spoon!", actual)
}

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	target := &strbuf.Target{Data: "!abbcF"}
	require.NoError(target.Analyze())

	t1 := &strbuf.RevTask{}
	t2 := &strbuf.Rot13Task{}

	ctx := context.TODO()
	target.Handle(ctx, t1)
	target.Handle(ctx, t2)
	require.Equal("Spoon!", target.Data)
}
