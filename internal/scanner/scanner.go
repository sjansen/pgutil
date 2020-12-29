package scanner

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Scanner provides a convenient interface for parsing SQL.
type Scanner struct {
	buf *bufio.Reader
}

// New returns a new Scanner.
func New(data string) (*Scanner, error) {
	if !utf8.ValidString(data) {
		return nil, ErrInvalidUTF8
	}

	r := strings.NewReader(data)
	return &Scanner{
		buf: bufio.NewReader(r),
	}, nil
}

// RequireWhitespace scans one or more whitespace runes,
// returning an error if none are found.
func (s *Scanner) RequireWhitespace() error {
	r, _, err := s.buf.ReadRune()
	switch {
	case err == io.EOF:
		return ErrWhitespaceExpected
	case err != nil:
		return err
	case unicode.IsSpace(r):
		return s.SkipWhitespace()
	}

	if err = s.buf.UnreadRune(); err != nil {
		return err
	}
	return ErrWhitespaceExpected
}

// SkipWhitespace scans zero or more whitespace runes.
func (s *Scanner) SkipWhitespace() error {
	for {
		r, _, err := s.buf.ReadRune()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case !unicode.IsSpace(r):
			return s.buf.UnreadRune()
		}
	}
}
