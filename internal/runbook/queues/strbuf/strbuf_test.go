package strbuf_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/queues/strbuf"
)

func TestRevTask(t *testing.T) {
	require := require.New(t)

	task := &strbuf.RevTask{}
	require.NoError(task.VerifyConfig())

	actual := task.Munge("!noopS")
	require.Equal("Spoon!", actual)
}

func TestRot13Task(t *testing.T) {
	require := require.New(t)

	task := &strbuf.Rot13Task{}
	require.NoError(task.VerifyConfig())

	actual := task.Munge("Fcbba!")
	require.Equal("Spoon!", actual)
}

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	queue := &strbuf.StrBuf{Message: "!abbcF"}
	require.NoError(queue.VerifyConfig())

	t1 := &strbuf.RevTask{}
	t2 := &strbuf.Rot13Task{}
	require.NoError(queue.VerifyTask(t1))
	require.NoError(queue.VerifyTask(t2))

	queue.Start(t1)
	queue.Start(t2)
	require.Equal("Spoon!", queue.Message)
}
