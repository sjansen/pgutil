package ddl

import (
	"fmt"
	"strings"
)

type parseError struct {
	cs   int
	data string
}

func (e *parseError) Error() string {
	cs := e.cs
	data := e.data
	if len(data) <= cs {
		return "error after end of data"
	}

OUTER:
	for {
		idx := strings.IndexRune(data, '\n')
		switch {
		case idx == -1:
			break OUTER
		case idx > cs:
			data = data[:idx]
			break OUTER
		default:
			idx += 1
			cs -= idx
			data = data[idx:]
		}
	}

	return fmt.Sprintf("%s\n%*s", data, cs+1, "^")
}
