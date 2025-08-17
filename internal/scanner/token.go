package scanner

import "strconv"

// TokenKind represents the type of token.
type TokenKind int

func (kind TokenKind) String() string {
	return strconv.Itoa(int(kind))
}

const (
	EOF TokenKind = iota // End of file

	// Operators
	Plus
	Minus
	Star
	Slash
	Percent
	LessThan
	GreaterThan
	Equal
	Not
	And
	Or
	Caret
	Question

	// Parentheses
	LeftParen
	RightParen
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace

	// Numbers
	Integer
	Float

	// Blanks and Whitespace
	Whitespace
	Newline
	Tab
	CarriageReturn

	// Punctuation
	Comma
	Semicolon
	Colon
	Dot
	SingleQuote
	DoubleQuote

	// Identifiers
	Identifier
)

type Token struct {
	Kind  TokenKind
	Value string
	Row   int
	Col   int
}
