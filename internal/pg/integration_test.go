// +build integration

package pg_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

func connect() (c *pg.Conn, err error) {
	options := &pg.Options{
		Log: logger.Discard(),
	}
	for retries := 0; retries < 5; retries++ {
		if c, err = pg.New(options); err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return
}

func TestConnectAndQuery(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	version, err := c.ServerVersion()
	require.NoError(err)
	require.NotEmpty(version)
}

func TestExec(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	query := `
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
`
	err = c.Exec(query)
	require.NoError(err)
}

func TestListColumns(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	columns, err := c.ListColumns("pg_catalog", "pg_class")
	require.NoError(err)
	require.NotEmpty(columns)
}

func TestListFunctions(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	functions, err := c.ListFunctions()
	require.NoError(err)
	require.NotEmpty(functions)
}

func TestListSchemas(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	schemas, err := c.ListSchemas()
	require.NoError(err)
	require.NotEmpty(schemas)
}

func TestListTables(t *testing.T) {
	require := require.New(t)

	c, err := connect()
	require.NoError(err)
	defer c.Close()

	tables, err := c.ListTables()
	require.NoError(err)
	require.NotEmpty(tables)
}
