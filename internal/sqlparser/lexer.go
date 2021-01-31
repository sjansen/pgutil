package sqlparser

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

const eof = 0

type lexer struct {
	result     interface{}
	err        string
	str        string
	mode       int
	currOffset int
	prevOffset int
	markHead   int
	markTail   int
	decoded    rune
}

func (l *lexer) Lex(lval *yySymType) int { // nolint: gocyclo
	if l.mode != 0 {
		tmp := l.mode
		l.mode = 0
		return tmp
	}

	l.markTail = l.prevOffset
	l.skipWS()

	ch := l.decoded
	switch {
	case ch == eof:
		return eof
	case ch >= '0' && ch <= '9':
		for l.decoded >= '0' && l.decoded <= '9' {
			l.decode()
		}
		return ICONST
	case ch == '\'':
		l.decode()
		for l.decoded != '\'' {
			l.decode()
		}
		l.decode()
		return SCONST
	case ch == '(' || ch == ')' || ch == '=' || ch == '.' || ch == ',' || ch == ';':
		l.decode()
		return int(ch)
	case ch == '+' || ch == '-' || ch == '*' || ch == '/':
		l.decode()
		return int(ch)
	case ch == '[' || ch == ']' || ch == '%' || ch == '^':
		l.decode()
		return int(ch)
	case ch == '~':
		l.decode()
		return Op
	case ch == '<':
		l.decode()
		switch l.decoded {
		case '=':
			l.decode()
			return LESS_EQUALS
		case '>':
			l.decode()
			return NOT_EQUALS
		default:
			return int(ch)
		}
	case ch == '>':
		l.decode()
		switch l.decoded {
		case '=':
			l.decode()
			return GREATER_EQUALS
		default:
			return int(ch)
		}
	case ch == '!':
		l.decode()
		switch l.decoded {
		case '=':
			l.decode()
			return NOT_EQUALS
		default:
			return int(ch)
		}
	case ch == ':':
		l.decode()
		switch l.decoded {
		case ':':
			l.decode()
			return TYPECAST
		default:
			return int(ch)
		}
	case isWordStart(ch):
		return l.scanWord(lval)
	}

	return UNEXPECTED_SYMBOL
}

// Error is called by the parser when there's a syntax error.
func (l *lexer) Error(s string) {
	l.err = s
}

func (l *lexer) decode() rune {
	r, size := utf8.DecodeRuneInString(l.str[l.currOffset:])
	if size == 0 {
		l.decoded = eof
		return eof
	}
	l.decoded = r
	l.prevOffset = l.currOffset
	l.currOffset += size
	if yyDebug > 6 {
		fmt.Printf("decoded: %q\n", r)
	}
	return r
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWordStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// a word is a keyword or an identifier
func (l *lexer) scanWord(lval *yySymType) int {
	buffer := bytes.Buffer{}
	ch := l.decoded
	for isWordStart(ch) || ch == '$' || isDigit(ch) {
		buffer.WriteRune(ch)
		ch = l.decode()
	}
	lower := string(bytes.ToLower(buffer.Bytes()))
	if keyword, ok := keywords[lower]; ok {
		return keyword
	}
	lval.str = lower
	return Identifier
}

func (l *lexer) setMark() {
	l.markHead = l.prevOffset
}

func (l *lexer) sinceMark() string {
	return l.str[l.markHead:l.markTail]
}

func (l *lexer) skipWS() {
	ch := l.decoded
	if ch == 0 {
		ch = l.decode()
	}
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		ch = l.decode()
	}
}
