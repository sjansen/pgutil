package sqlparser_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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
	parseAndCompare(t, "testdata/statements/*.sql", func(str string) (interface{}, error) {
		return sqlparser.Parse(str)
	})
}

func TestParseCheck(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/fragments/check*.sql", func(str string) (interface{}, error) {
		return sqlparser.ParseCheck(str)
	})
}

func TestParseForeignKey(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/fragments/foreign_key*.sql", func(str string) (interface{}, error) {
		return sqlparser.ParseForeignKey(str)
	})
}

func TestParseCreateIndex(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/statements/create_index*.sql", func(str string) (interface{}, error) {
		return sqlparser.ParseCreateIndex(str)
	})
}

func TestParseCreateTrigger(t *testing.T) {
	t.Parallel()
	// sqlparser.EnableDebugLogging()
	parseAndCompare(t, "testdata/statements/create_trigger*.sql", func(str string) (interface{}, error) {
		return sqlparser.ParseCreateTrigger(str)
	})
}

func parseAndCompare(t *testing.T, pattern string, fn func(string) (interface{}, error)) {
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

			data, err := ioutil.ReadFile(path)
			require.NoError(err)

			var actual interface{}
			sql := strings.TrimSpace(string(data))
			actual, err = fn(sql)
			require.NoError(err)
			require.NotNil(actual)

			data, err = ioutil.ReadFile(
				replaceExtension(path, "json"),
			)
			require.NoError(err)

			expectedResult := &Result{}
			err = json.Unmarshal(data, expectedResult)
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
