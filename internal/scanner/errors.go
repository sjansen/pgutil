package scanner

import "errors"

// ErrInvalidUTF8 is returned for strings containing non-UTF8 bytes.
var ErrInvalidUTF8 = errors.New("invalid utf8")

// ErrWhitespaceExpected when a whitespace rune is expected but not found.
var ErrWhitespaceExpected = errors.New("whitespace expected")
