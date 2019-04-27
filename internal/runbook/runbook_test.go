package runbook_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/strbuf"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	p := parser.Parser{
		Targets: map[string]types.TargetFactory{
			"strbuf": &strbuf.TargetFactory{},
		},
	}

	expected := &types.Runbook{
		Targets: types.Targets{
			"strbuf": &strbuf.Target{
				Data: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Tasks: types.Tasks{
			"encrypted": {
				Target: "strbuf",
				Config: &strbuf.Echo{},
			},
			"decrypted": {
				After:  []string{"reverse", "rotate"},
				Target: "strbuf",
				Config: &strbuf.Echo{},
			},
			"reverse": {
				After:  []string{"encrypted"},
				Target: "strbuf",
				Config: &strbuf.Rev{},
			},
			"rotate": {
				After:  []string{"encrypted"},
				Target: "strbuf",
				Config: &strbuf.Rot13{},
			},
		},
	}

	actual, err := p.Parse("testdata/message.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)
}

func TestRun(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	err := runbook.Run("testdata/message.jsonnet", &stdout, &stderr)
	require.NoError(err)

	expected := `.ravgyniB ehbl xaveq bg rehf rO
Be sure to drink your Ovaltine.
`
	require.Equal(expected, stdout.String())
	require.Equal("", stderr.String())
}
