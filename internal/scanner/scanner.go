package scanner

import (
	"unicode"
	"unicode/utf8"
)

// Scanner provides a convenient interface for parsing SQL.
type Scanner struct {
	str string
	idx int

	currLine  int
	nextLine  int
	prevIndex int
}

// New returns a new Scanner.
func New(s string) (*Scanner, error) {
	if !utf8.ValidString(s) {
		return nil, ErrInvalidUTF8
	}
	return &Scanner{
		str:       s,
		nextLine:  -1,
		prevIndex: -1,
	}, nil
}

func (s *Scanner) readRune() rune {
	s.prevIndex = s.idx
	r, size := utf8.DecodeRuneInString(s.str[s.idx:])
	s.idx += size
	switch {
	case r == '\r':
		tmp, size := utf8.DecodeRuneInString(s.str[s.idx:])
		if tmp == '\n' {
			s.idx += size
			r = tmp
		}
		fallthrough
	case r == '\n':
		s.nextLine = s.idx + 1
	case s.nextLine > 0:
		s.currLine = s.nextLine
		s.nextLine = -1
		s.prevIndex = -1
	}
	return r
}

func (s *Scanner) unreadRune() {
	if s.prevIndex >= 0 {
		s.idx = s.prevIndex
	}
}

// RequireKeyword consumes runes if they match after being uppercased.
func (s *Scanner) RequireKeyword(keyword string) error {
	for _, expected := range keyword {
		actual := s.readRune()
		if expected != unicode.ToUpper(actual) {
			return ErrKeywordExpected
		}
	}
	return nil
}

// RequireWhitespace consumes one or more whitespace runes.
func (s *Scanner) RequireWhitespace() error {
	r := s.readRune()
	if unicode.IsSpace(r) {
		return s.SkipWhitespace()
	}
	return ErrWhitespaceExpected
}

// SkipWhitespace consumes zero or more whitespace runes.
func (s *Scanner) SkipWhitespace() error {
	r := s.readRune()
	for unicode.IsSpace(r) {
		r = s.readRune()
	}
	s.unreadRune()
	return nil
}
