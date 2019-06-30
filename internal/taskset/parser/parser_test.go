package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/taskset/parser"
	"github.com/sjansen/pgutil/internal/taskset/pg"
	"github.com/sjansen/pgutil/internal/taskset/sh"
	"github.com/sjansen/pgutil/internal/taskset/types"
)

func newParser() *parser.Parser {
	log := logger.Discard()
	return &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pg.TargetFactory{
				Log: log,
			},
			"sh": &sh.TargetFactory{
				Log: log,
			},
		},
	}

}

func TestParse(t *testing.T) {
	require := require.New(t)

	p := newParser()
	_, err := p.Parse("testdata/simple.hcl")
	require.NoError(err)
}
