package scanner

// TokenKind represents the type of token.
type TokenKind int

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
	kind  TokenKind
	value string
	row   int
	col   int
}
