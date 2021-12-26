//go:build docker && integration
// +build docker,integration

package main_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/pg"
	"github.com/sjansen/pgutil/internal/testutil"
)

const actualPath = "testdata/actual.hcl"
const expectedPath = "testdata/expected.hcl"

func TestConnectAndQuery(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			assert := assert.New(t)
			require := require.New(t)

			expected, err := ioutil.ReadFile(expectedPath)
			require.NoError(err)

			ctx := context.TODO()
			c, err := testutil.Connect(ctx)
			require.NoError(err)
			defer c.Close(ctx)

			db, err := c.InspectDatabase(ctx, &pg.InspectOptions{
				SortColumns: true,
				SortIndexes: true,
			})
			require.NoError(err)

			m := &ddl.DatabaseMetadata{
				Host:          "db.example.com",
				Database:      "example",
				ServerVersion: "42",
				Timestamp:     time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC),
			}

			buf := &bytes.Buffer{}
			err = db.WriteHCL(buf, m)
			require.NoError(err)

			actual := buf.Bytes()
			if !assert.Equal(string(expected), string(actual)) {
				ioutil.WriteFile(actualPath, actual, 0666)
				t.Log(
					"Temp JSON file created to facilitate debugging.",
					"\nexpected:", expectedPath,
					"\nactual:", actualPath,
				)
			}
		})
	}
}
