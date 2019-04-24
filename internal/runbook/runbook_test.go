package runbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/strbuf"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Targets: map[string]types.TargetFactory{
			"strbuf": &strbuf.TargetFactory{},
		},
	}

	expected := &types.Runbook{
		Targets: map[string]types.Target{
			"strbuf": &strbuf.Target{
				Data: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Tasks: map[string]*types.Task{
			"encrypted": {
				Target: "strbuf",
				Config: &strbuf.EchoTask{},
			},
			"decrypted": {
				After:  []string{"reverse", "rotate"},
				Target: "strbuf",
				Config: &strbuf.EchoTask{},
			},
			"reverse": {
				After:  []string{"encrypted"},
				Target: "strbuf",
				Config: &strbuf.RevTask{},
			},
			"rotate": {
				After:  []string{"encrypted"},
				Target: "strbuf",
				Config: &strbuf.Rot13Task{},
			},
		},
	}

	actual, err := l.Parse("testdata/message.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)
}
