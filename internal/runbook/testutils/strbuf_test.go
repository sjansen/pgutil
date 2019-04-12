package testutils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/testutils"
)

func TestRevTask(t *testing.T) {
	require := require.New(t)

	task := &testutils.RevTask{}
	require.NoError(task.VerifyConfig())

	actual := task.Munge("!noopS")
	require.Equal("Spoon!", actual)
}

func TestRot13Task(t *testing.T) {
	require := require.New(t)

	task := &testutils.Rot13Task{}
	require.NoError(task.VerifyConfig())

	actual := task.Munge("Fcbba!")
	require.Equal("Spoon!", actual)
}

func TestStrBuf(t *testing.T) {
	require := require.New(t)

	queue := &testutils.StrBuf{Message: "!abbcF"}
	require.NoError(queue.VerifyConfig())

	t1 := &testutils.RevTask{}
	t2 := &testutils.Rot13Task{}
	require.NoError(queue.VerifyTask(t1))
	require.NoError(queue.VerifyTask(t2))

	queue.Start(t1)
	queue.Start(t2)
	require.Equal("Spoon!", queue.Message)
}
