package scanner

type Scanner struct {
	readPosition int
	position     int
	ch           byte
	input        string
	row          int
	col          int
}

func (sc *Scanner) NextToken() Token {

	switch sc.ch {

	// Operators
	case '+':
		return sc.readSimpleToken(Plus)
	case '-':
		return sc.readSimpleToken(Minus)
	case '*':
		return sc.readSimpleToken(Star)
	case '/':
		return sc.readSimpleToken(Slash)
	case '%':
		return sc.readSimpleToken(Percent)
	case '<':
		return sc.readSimpleToken(LessThan)
	case '>':
		return sc.readSimpleToken(GreaterThan)
	case '=':
		return sc.readSimpleToken(Equal)
	case '!':
		return sc.readSimpleToken(Not)
	case '&':
		return sc.readSimpleToken(And)
	case '|':
		return sc.readSimpleToken(Or)
	case '^':
		return sc.readSimpleToken(Caret)
	case '?':
		return sc.readSimpleToken(Question)

	// Parentheses
	case '(':
		return sc.readSimpleToken(LeftParen)
	case ')':
		return sc.readSimpleToken(RightParen)
	case '[':
		return sc.readSimpleToken(LeftBracket)
	case ']':
		return sc.readSimpleToken(RightBracket)
	case '{':
		return sc.readSimpleToken(LeftBrace)
	case '}':
		return sc.readSimpleToken(RightBrace)

	// Punctuation
	case ',':
		return sc.readSimpleToken(Comma)
	case ';':
		return sc.readSimpleToken(Semicolon)
	case ':':
		return sc.readSimpleToken(Colon)
	case '.':
		return sc.readSimpleToken(Dot)
	case '\'':
		return sc.readSimpleToken(SingleQuote)
	case '"':
		return sc.readSimpleToken(DoubleQuote)

	// Whitespaces
	case ' ':
		return sc.readSimpleToken(Whitespace)
	case '\t':
		return sc.readSimpleToken(Tab)
	case '\n':
		return sc.readSimpleToken(Newline)
	case '\r':
		return sc.readSimpleToken(CarriageReturn)

	// EOF
	case 0:
		return Token{Kind: EOF, Value: "", Row: sc.row, Col: sc.col}

	default:
		if sc.isDigit(sc.ch) {
			return sc.readNumber()
		}
		return sc.readIdentifier()
	}
}

func (sc *Scanner) readChar() byte {
	if sc.readPosition >= len(sc.input) {
		sc.ch = 0
	} else {
		sc.ch = sc.input[sc.readPosition]
	}
	if sc.ch == '\n' {
		sc.row++
		sc.col = 0
	} else {
		sc.col++
	}
	sc.position = sc.readPosition
	sc.readPosition++
	return sc.ch
}

func (sc *Scanner) readSimpleToken(kind TokenKind) Token {
	token := Token{Kind: kind, Value: string(sc.ch), Row: sc.row, Col: sc.col}
	sc.readChar()
	return token
}

func (sc *Scanner) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (sc *Scanner) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (sc *Scanner) readNumber() Token {
	tok := Token{Row: sc.row, Col: sc.col, Kind: Integer}

	for sc.isDigit(sc.ch) || sc.ch == '.' {
		if sc.ch == '.' {
			tok.Kind = Float
		}
		tok.Value += string(sc.ch)
		sc.readChar()
	}

	return tok
}

func (sc *Scanner) readIdentifier() Token {
	tok := Token{Row: sc.row, Col: sc.col, Kind: Identifier}

	for sc.isLetter(sc.ch) || sc.isDigit(sc.ch) || sc.ch == '_' {
		tok.Value += string(sc.ch)
		sc.readChar()
	}

	return tok
}

func (sc *Scanner) readNewline() Token {
	tok := Token{Kind: Newline, Value: "\n", Row: sc.row, Col: sc.col}
	sc.row++
	sc.col = 0
	sc.readChar()
	return tok
}

func NewScanner(input string) *Scanner {
	sc := &Scanner{
		input:        input,
		readPosition: 0,
		position:     0,
		ch:           0,
		row:          1,
		col:          0,
	}

	sc.readChar()
	return sc
}
