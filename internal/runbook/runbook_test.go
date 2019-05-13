package runbook_test

import (
	"bytes"
	"testing"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/strbuf"
	"github.com/sjansen/pgutil/internal/runbook/types"
	"github.com/sjansen/pgutil/internal/sys"
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

func TestRunWithFailedTask(t *testing.T) {
	require := require.New(t)
	defer leaktest.Check(t)()

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}

	filename := "testdata/fail.jsonnet"
	err := runbook.Run(sys, filename)
	require.Error(err, filename)

	expected := "foo\nbar\nbaz\n"
	actual := stdout.String()
	require.Equal(expected, actual)
}
