package runbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/testutils"
)

func TestLoad(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Queues: map[string]func() parser.Queue{
			"strbuf": func() parser.Queue { return &testutils.StrBuf{} },
		},
		Tasks: map[string]func() parser.TaskConfig{
			"strbuf/echo":  func() parser.TaskConfig { return &testutils.EchoTask{} },
			"strbuf/rev":   func() parser.TaskConfig { return &testutils.RevTask{} },
			"strbuf/rot13": func() parser.TaskConfig { return &testutils.Rot13Task{} },
		},
	}

	expected := &parser.Runbook{
		Queues: map[string]parser.Queue{
			"strbuf": &testutils.StrBuf{
				Message: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Tasks: map[string]*parser.Task{
			"encrypted": {
				Queue:  "strbuf",
				Config: &testutils.EchoTask{},
			},
			"decrypted": {
				After:  []string{"reverse", "rotate"},
				Queue:  "strbuf",
				Config: &testutils.EchoTask{},
			},
			"reverse": {
				After:  []string{"encrypted"},
				Queue:  "strbuf",
				Config: &testutils.RevTask{},
			},
			"rotate": {
				After:  []string{"encrypted"},
				Queue:  "strbuf",
				Config: &testutils.Rot13Task{},
			},
		},
	}

	actual, err := l.Load("testdata/message.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)
}
