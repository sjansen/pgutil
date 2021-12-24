//go:build integration
// +build integration

package pg_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/pg"
	"github.com/sjansen/pgutil/internal/testutil"
)

func TestConnectAndQuery(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			version, err := c.ServerVersion(ctx)
			require.NoError(err)
			require.NotEmpty(version)
		})
	}
}

func TestDescribeFunction(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			function, err := c.DescribeFunction(ctx, "pg_catalog", "current_database")
			require.NoError(err)
			require.NotNil(function)

			function, err = c.DescribeFunction(ctx, "public", "no_such_function")
			require.Equal(pg.ErrNotFound, err)
			require.Nil(function)
		})
	}
}

func TestDescribeTrigger(t *testing.T) {
	if os.Getenv("PGUTIL_TEST_TAGS") == "" {
		t.Skip("missing $PGUTIL_TEST_TAGS")
	}
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			trigger, err := c.DescribeTrigger(ctx, "public", "measurement", "update_modified_column")
			require.NoError(err)
			require.NotNil(trigger)

			trigger, err = c.DescribeTrigger(ctx, "public", "no_such_table", "no_such_trigger")
			require.Equal(pg.ErrNotFound, err)
			require.Nil(trigger)
		})
	}
}

func TestExec(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			query := `
BEGIN
;
CREATE TABLE IF NOT EXISTS measurements (
    id BIGSERIAL NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    value DOUBLE PRECISION NOT NULL
)
;
DELETE FROM measurements
WHERE timestamp < now() - interval '5 minutes'
;
INSERT INTO measurements
    (timestamp, value)
VALUES
    (now(), random())
;
INSERT INTO measurements
    (timestamp, value)
VALUES
    (now(), random())
;
ROLLBACK
;
`
			err = c.Exec(ctx, query)
			require.NoError(err)
		})
	}
}

func TestListColumns(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			columns, err := c.ListColumns(ctx, "pg_catalog", "pg_class")
			require.NoError(err)
			require.NotEmpty(columns)
		})
	}
}

func TestListFunctions(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			functions, err := c.ListFunctions(ctx)
			require.NoError(err)
			require.NotEmpty(functions)
		})
	}
}

func TestListSchemas(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			schemas, err := c.ListSchemas(ctx)
			require.NoError(err)
			require.NotEmpty(schemas)
		})
	}
}

func TestListSequences(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			sequences, err := c.ListSequences(ctx)
			require.NoError(err)
			require.NotEmpty(sequences)
		})
	}
}

func TestListTables(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			ctx := context.Background()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			tables, err := c.ListTables(ctx)
			require.NoError(err)
			require.NotEmpty(tables)
		})
	}
}
