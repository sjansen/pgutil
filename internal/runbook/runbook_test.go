package runbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/queues/strbuf"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Queues: map[string]func() parser.Queue{
			"strbuf": func() parser.Queue { return &strbuf.StrBuf{} },
		},
		Tasks: map[string]func() parser.Task{
			"strbuf/echo":  func() parser.Task { return &strbuf.EchoTask{} },
			"strbuf/rev":   func() parser.Task { return &strbuf.RevTask{} },
			"strbuf/rot13": func() parser.Task { return &strbuf.Rot13Task{} },
		},
	}

	expected := &parser.Runbook{
		Queues: map[string]parser.Queue{
			"strbuf": &strbuf.StrBuf{
				Message: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Steps: map[string]*parser.Step{
			"encrypted": {
				Queue: "strbuf",
				Task:  &strbuf.EchoTask{},
			},
			"decrypted": {
				After: []string{"reverse", "rotate"},
				Queue: "strbuf",
				Task:  &strbuf.EchoTask{},
			},
			"reverse": {
				After: []string{"encrypted"},
				Queue: "strbuf",
				Task:  &strbuf.RevTask{},
			},
			"rotate": {
				After: []string{"encrypted"},
				Queue: "strbuf",
				Task:  &strbuf.Rot13Task{},
			},
		},
	}

	actual, err := l.Parse("testdata/message.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)
}
