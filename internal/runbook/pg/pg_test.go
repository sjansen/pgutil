package pg_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook/pg"
)

func TestTasks(t *testing.T) {
	require := require.New(t)

	factory := &pg.TargetFactory{
		Log: logger.Discard(),
	}
	target := factory.NewTarget()
	err := target.Start()
	require.NoError(err)

	ctx := context.TODO()
	require.NoError(target.Analyze())

	task := &pg.Exec{}
	require.Error(task.Check())

	for _, tc := range []struct {
		sql string
		err bool
	}{{
		sql: "SELECT now()",
		err: false,
	}, {
		sql: "SHOW search_path",
		err: false,
	}, {
		sql: "Show me the money!",
		err: true,
	}} {
		task := &pg.Exec{SQL: tc.sql}
		require.NoError(task.Check())

		err := target.Handle(ctx, task)
		if tc.err {
			require.Error(err, tc.sql)
		} else {
			require.NoError(err, tc.sql)
		}
	}
}
