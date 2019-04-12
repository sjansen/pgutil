package runbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/loader"
	"github.com/sjansen/pgutil/internal/runbook/testutils"
)

func TestLoad(t *testing.T) {
	require := require.New(t)

	l := loader.Loader{
		Queues: map[string]func() loader.Queue{
			"strbuf": func() loader.Queue { return &testutils.StrBuf{} },
		},
		Tasks: map[string]func() loader.TaskConfig{
			"strbuf/echo":  func() loader.TaskConfig { return &testutils.EchoTask{} },
			"strbuf/rev":   func() loader.TaskConfig { return &testutils.RevTask{} },
			"strbuf/rot13": func() loader.TaskConfig { return &testutils.Rot13Task{} },
		},
	}

	expected := &loader.Runbook{
		Queues: map[string]loader.Queue{
			"strbuf": &testutils.StrBuf{
				Message: ".ravgyniB ehbl xaveq bg rehf rO",
			},
		},
		Tasks: map[string]*loader.Task{
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
