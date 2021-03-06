package scanner

import "errors"

// ErrInvalidUTF8 is returned for strings containing non-UTF8 bytes.
var ErrInvalidUTF8 = errors.New("invalid utf8")

// ErrKeywordExpected when a keyword is expected but not found.
var ErrKeywordExpected = errors.New("keyword expected")

// ErrWhitespaceExpected when a whitespace rune is expected but not found.
var ErrWhitespaceExpected = errors.New("whitespace expected")
