package commands_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/mocks"
)

func TestRunbook(t *testing.T) {
	require := require.New(t)

	db := new(mocks.DB)
	db.On("Close").Return(nil)
	db.On("Exec", mock.AnythingOfType("string")).Return(nil)
	processes := make([]*mocks.Process, 0)
	deps := &commands.Dependencies{
		DB: func(opts map[string]string) (commands.DB, error) {
			return db, nil
		},
		Process: func(args []string) commands.Process {
			p := &mocks.Process{}
			p.On("Run", mock.Anything, mock.Anything).Return(nil)
			if len(args) > 1 && args[0] == "echo" {
				p.Stdout = strings.Join(args[1:], " ") + "\n"
			}
			processes = append(processes, p)
			return p
		},
	}

	var stdout, stderr bytes.Buffer
	cmd := &commands.RunBookRunCmd{
		File: "testdata/simple.jsonnet",
	}
	err := cmd.Run(&stdout, &stderr, deps)
	require.NoError(err)
	require.NotEmpty(stdout.String()) // TODO
	require.Empty(stderr.String())
	db.AssertExpectations(t)
	for _, p := range processes {
		p.AssertExpectations(t)
	}
}
