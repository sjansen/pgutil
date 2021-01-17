package sqlparser_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/sqlparser"
)

type Result struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func replaceExtension(path, replacement string) string {
	ext := filepath.Ext(path)
	return path[0:len(path)-len(ext)] + replacement
}

func TestParser(t *testing.T) {
	t.Parallel()
	sqlparser.EnableDebugLogging()

	testcases, err := filepath.Glob("testdata/*.sql")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for _, path := range testcases {
		path := path
		basename := filepath.Base(path)
		t.Run(basename, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			sql, err := ioutil.ReadFile(path)
			require.NoError(err)

			actualStmt, err := sqlparser.Parse(sql)
			require.NoError(err)
			require.NotNil(actualStmt)

			tmp, err := ioutil.ReadFile(
				replaceExtension(path, ".json"),
			)
			require.NoError(err)

			expected := &Result{}
			err = json.Unmarshal(tmp, expected)
			require.NoError(err)

			actualType := reflect.TypeOf(actualStmt)
			require.Equal(expected.Type, actualType.Elem().String())

			expectedStmt := reflect.New(actualType.Elem()).Interface()
			err = json.Unmarshal(expected.Value, expectedStmt)
			require.NoError(err)

			if !assert.Equal(expectedStmt, actualStmt) {
				actual := &Result{
					Type: actualType.Name(),
				}
				marshalled, err := json.MarshalIndent(actual, "", "  ")
				require.NoError(err)

				f, err := ioutil.TempFile("", "actual.*.json")
				require.NoError(err)
				defer f.Close()

				if _, err := f.Write(marshalled); err == nil {
					t.Log(
						"Temp JSON file created to facilitate debugging.",
						"\nexpected:", path,
						"\nactual:", f.Name(),
					)
				}
			}
		})
	}
}
