package scanner

import (
	"unicode"
	"unicode/utf8"
)

// Position represent the position of a rune, starting at (1,1).
type Position struct {
	Line   int
	Column int
}

// State represents the state of a scanner.
type State struct {
	position Position

	nextRuneOffset int  //nolint:structcheck
	prevRune       rune //nolint:structcheck
}

// Scanner provides a convenient interface for parsing SQL.
type Scanner struct {
	data string
	State
}

// New returns a new Scanner.
func New(s string) (*Scanner, error) {
	if !utf8.ValidString(s) {
		return nil, ErrInvalidUTF8
	}
	return &Scanner{
		data: s,
		State: State{
			position: Position{
				Line: 1,
			},
		},
	}, nil
}

func (s *Scanner) readRune() rune {
	if s.nextRuneOffset >= len(s.data) {
		return utf8.RuneError
	}
	r, size := utf8.DecodeRuneInString(s.data[s.nextRuneOffset:])
	s.nextRuneOffset += size
	switch {
	case r == '\r':
		r = '\n'
		tmp, size := utf8.DecodeRuneInString(s.data[s.nextRuneOffset:])
		if tmp == '\n' {
			s.nextRuneOffset += size
		}
		fallthrough
	case r == '\n':
		s.position.Column++
	case s.prevRune == '\n':
		s.position.Line++
		s.position.Column = 1
	default:
		s.position.Column++
	}
	s.prevRune = r
	return r
}

// Position returns the line & column of the most recently read rune,
// starting at (1, 1).
func (s *Scanner) Position() Position {
	return s.position
}

// Reset restores the scanner to a previous state.
func (s *Scanner) Reset(snapshot State) {
	s.State = snapshot
}

// Snapshot returns a snapshot of the scanner's state.
func (s *Scanner) Snapshot() State {
	return s.State
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
	switch s.readRune() {
	case ' ', '\t', '\r', '\n':
		return s.SkipWhitespace()
	}
	return ErrWhitespaceExpected
}

// SkipWhitespace consumes zero or more whitespace runes.
func (s *Scanner) SkipWhitespace() error {
	for s.nextRuneOffset < len(s.data) {
		switch s.data[s.nextRuneOffset] {
		case ' ', '\t', '\r', '\n':
			s.readRune()
		default:
			return nil
		}
	}
	return nil
}
