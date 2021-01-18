package sqlparser_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

func TestParse(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/statements/*.sql", func(buf []byte) (interface{}, error) {
		return sqlparser.Parse(buf)
	})
}

func TestParseForeignKey(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/fragments/foreign_key*.sql", func(buf []byte) (interface{}, error) {
		return sqlparser.ParseForeignKey(buf)
	})
}

func parseAndCompare(t *testing.T, pattern string, fn func([]byte) (interface{}, error)) {
	testcases, err := filepath.Glob(pattern)
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

			var actual interface{}
			actual, err = fn(sql)
			require.NoError(err)
			require.NotNil(actual)

			buf, err := ioutil.ReadFile(
				replaceExtension(path, "json"),
			)
			require.NoError(err)

			expectedResult := &Result{}
			err = json.Unmarshal(buf, expectedResult)
			require.NoError(err)

			typ := reflect.TypeOf(actual)
			require.Equal(expectedResult.Type, typ.Elem().String())

			expected := reflect.New(typ.Elem()).Interface()
			err = json.Unmarshal(expectedResult.Value, expected)
			require.NoError(err)

			if !assert.Equal(expected, actual) {
				tmp, err := json.MarshalIndent(actual, "", "  ")
				require.NoError(err)

				result := &Result{Type: typ.Elem().String()}
				err = result.Value.UnmarshalJSON(tmp)
				require.NoError(err)

				marshalled, err := json.MarshalIndent(result, "", "  ")
				require.NoError(err)

				f, err := os.OpenFile(replaceExtension(path, "actual"),
					os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755,
				)
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

func replaceExtension(path, replacement string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return path + "." + replacement
	}
	return path[0:len(path)-len(ext)] + "." + replacement
}
