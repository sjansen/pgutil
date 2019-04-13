package runbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/testutils"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Queues: map[string]func() parser.Queue{
			"strbuf": func() parser.Queue { return &testutils.StrBuf{} },
		},
		Tasks: map[string]func() parser.Task{
			"strbuf/echo":  func() parser.Task { return &testutils.EchoTask{} },
			"strbuf/rev":   func() parser.Task { return &testutils.RevTask{} },
			"strbuf/rot13": func() parser.Task { return &testutils.Rot13Task{} },
		},
	}

	expected := &parser.Runbook{
		Queues: map[string]parser.Queue{
			"strbuf": &testutils.StrBuf{
				Message: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Steps: map[string]*parser.Step{
			"encrypted": {
				Queue: "strbuf",
				Task:  &testutils.EchoTask{},
			},
			"decrypted": {
				After: []string{"reverse", "rotate"},
				Queue: "strbuf",
				Task:  &testutils.EchoTask{},
			},
			"reverse": {
				After: []string{"encrypted"},
				Queue: "strbuf",
				Task:  &testutils.RevTask{},
			},
			"rotate": {
				After: []string{"encrypted"},
				Queue: "strbuf",
				Task:  &testutils.Rot13Task{},
			},
		},
	}

	actual, err := l.Parse("testdata/message.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)
}
