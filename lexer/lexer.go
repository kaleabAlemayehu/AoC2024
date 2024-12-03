package lexer

import (
	"strconv"
)

type Lexer struct {
	input        []byte
	position     int  // Current position of a character that is being read
	readPosition int  // The next character position
	ch           byte // Current character that is being read
}

func NewLexer(input []byte) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() int {
	var tok int
	l.skipWhitespace()
	if isDigit(l.ch) {
		tok, err := strconv.Atoi(string(l.readNumber()))
		if err != nil {
			panic(err)
		}
		return tok
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNumber() []byte {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
