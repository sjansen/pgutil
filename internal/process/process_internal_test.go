package process

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateArg0(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Arg0  string
		Legal bool
	}{{
		Arg0: "echo", Legal: true,
	}, {
		Arg0: "scripts/do-something-magic", Legal: false,
	}, {
		Arg0: "./scripts/do-something-magic", Legal: true,
	}, {
		Arg0: "../scripts/do-something-magic", Legal: true,
	}, {
		Arg0: "/sbin/halt", Legal: true,
	}} {
		err := validateArg0(tc.Arg0)
		if tc.Legal {
			require.NoError(err, tc.Arg0)
		} else {
			require.Error(err, tc.Arg0)
		}
	}
}
