package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

const eof = 0

type Lexer struct {
	Statement Statement
	buf       []byte
	decoded   rune
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.skipWS()

	ch := l.decoded
	switch {
	case ch == eof:
		return eof
	case ch == ',' || ch == ';':
		l.decode()
		return int(ch)
	case isWordStart(ch):
		return l.scanWord(lval)
	}

	return LEX_ERROR
}

func (l *Lexer) Error(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func (l *Lexer) decode() rune {
	r, size := utf8.DecodeRune(l.buf)
	if size == 0 {
		l.decoded = eof
		return eof
	}
	l.decoded = r
	l.buf = l.buf[size:]
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

func (l *Lexer) skipWS() {
	ch := l.decoded
	if ch == 0 {
		ch = l.decode()
	}
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		ch = l.decode()
	}
}
