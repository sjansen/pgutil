package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type Process struct {
	mock.Mock
	Stdout string
	Stderr string
}

func (p *Process) Run(stdout, stderr io.Writer) error {
	result := p.Called(stdout, stderr)
	stdout.Write([]byte(p.Stdout))
	stderr.Write([]byte(p.Stderr))
	return result.Error(0)
}
