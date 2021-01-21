package sqlparser

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

const eof = 0

// Lexer tokenizes SQL statments for the parser.
type Lexer struct {
	result  interface{}
	err     string
	buf     []byte
	mode    int
	offset  int
	prev    int
	mark    int
	decoded rune
}

// Lex returns the next token for the parser.
func (l *Lexer) Lex(lval *yySymType) int { // nolint: gocyclo
	if l.mode != 0 {
		tmp := l.mode
		l.mode = 0
		return tmp
	}

	l.skipWS()

	ch := l.decoded
	switch {
	case ch == eof:
		return eof
	case ch == '(' || ch == ')' || ch == '=' || ch == '.' || ch == ',' || ch == ';':
		l.decode()
		return int(ch)
	case ch == '<' || ch == '>' || ch == '+' || ch == '-' || ch == '*' || ch == '/':
		l.decode()
		return int(ch)
	case ch == '[' || ch == ']' || ch == ':' || ch == '%' || ch == '^':
		l.decode()
		return int(ch)
	case isWordStart(ch):
		return l.scanWord(lval)
	}

	return UNEXPECTED_SYMBOL
}

// Error is called by the parser when there's a syntax error.
func (l *Lexer) Error(s string) {
	l.err = s
}

func (l *Lexer) decode() rune {
	r, size := utf8.DecodeRune(l.buf[l.offset:])
	if size == 0 {
		l.decoded = eof
		return eof
	}
	l.decoded = r
	l.prev = l.offset
	l.offset += size
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
func (l *Lexer) scanWord(lval *yySymType) int {
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

func (l *Lexer) setMark() {
	l.mark = l.prev
}

func (l *Lexer) sinceMark() []byte {
	// TODO: eliminate this hack
	_, size := utf8.DecodeLastRune(l.buf[:l.prev])
	return l.buf[l.mark : l.prev-size]
}

func (l *Lexer) skipWS() {
	ch := l.decoded
	if ch == 0 {
		ch = l.decode()
	}
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		ch = l.decode()
	}
}
