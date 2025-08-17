package scanner

import "testing"

func TestScanner(t *testing.T) {

	testCases := []struct {
		title    string
		input    string
		expected []Token
	}{
		{
			"empty_input",
			"",
			[]Token{
				{Kind: EOF, Value: "", Row: 1, Col: 1},
			},
		},
		{
			"operators",
			"+-*/%<>=!&|^?",
			[]Token{
				{Kind: Plus, Value: "+", Row: 1, Col: 1},
				{Kind: Minus, Value: "-", Row: 1, Col: 2},
				{Kind: Star, Value: "*", Row: 1, Col: 3},
				{Kind: Slash, Value: "/", Row: 1, Col: 4},
				{Kind: Percent, Value: "%", Row: 1, Col: 5},
				{Kind: LessThan, Value: "<", Row: 1, Col: 6},
				{Kind: GreaterThan, Value: ">", Row: 1, Col: 7},
				{Kind: Equal, Value: "=", Row: 1, Col: 8},
				{Kind: Not, Value: "!", Row: 1, Col: 9},
				{Kind: And, Value: "&", Row: 1, Col: 10},
				{Kind: Or, Value: "|", Row: 1, Col: 11},
				{Kind: Caret, Value: "^", Row: 1, Col: 12},
				{Kind: Question, Value: "?", Row: 1, Col: 13},
				{Kind: EOF, Value: "", Row: 1, Col: 14},
			},
		},
		{
			"numbers",
			"123 456.789",
			[]Token{
				{Kind: Integer, Value: "123", Row: 1, Col: 1},
				{Kind: Whitespace, Value: " ", Row: 1, Col: 4},
				{Kind: Float, Value: "456.789", Row: 1, Col: 5},
				{Kind: EOF, Value: "", Row: 1, Col: 12},
			},
		},
		{
			"identifiers",
			"var1 var_2 var3",
			[]Token{
				{Kind: Identifier, Value: "var1", Row: 1, Col: 1},
				{Kind: Whitespace, Value: " ", Row: 1, Col: 5},
				{Kind: Identifier, Value: "var_2", Row: 1, Col: 6},
				{Kind: Whitespace, Value: " ", Row: 1, Col: 11},
				{Kind: Identifier, Value: "var3", Row: 1, Col: 12},
				{Kind: EOF, Value: "", Row: 1, Col: 16},
			},
		},
		{
			"parentheses",
			"()[]{}",
			[]Token{
				{Kind: LeftParen, Value: "(", Row: 1, Col: 1},
				{Kind: RightParen, Value: ")", Row: 1, Col: 2},
				{Kind: LeftBracket, Value: "[", Row: 1, Col: 3},
				{Kind: RightBracket, Value: "]", Row: 1, Col: 4},
				{Kind: LeftBrace, Value: "{", Row: 1, Col: 5},
				{Kind: RightBrace, Value: "}", Row: 1, Col: 6},
				{Kind: EOF, Value: "", Row: 1, Col: 7},
			},
		},
		{
			"punctuation",
			",.;:\"'",
			[]Token{
				{Kind: Comma, Value: ",", Row: 1, Col: 1},
				{Kind: Dot, Value: ".", Row: 1, Col: 2},
				{Kind: Semicolon, Value: ";", Row: 1, Col: 3},
				{Kind: Colon, Value: ":", Row: 1, Col: 4},
				{Kind: DoubleQuote, Value: "\"", Row: 1, Col: 5},
				{Kind: SingleQuote, Value: "'", Row: 1, Col: 6},
				{Kind: EOF, Value: "", Row: 1, Col: 7},
			},
		},
		{
			"whitespace_and_newline",
			" \t\nabc\r",
			[]Token{
				{Kind: Whitespace, Value: " ", Row: 1, Col: 1},
				{Kind: Tab, Value: "\t", Row: 1, Col: 2},
				{Kind: Newline, Value: "\n", Row: 2, Col: 0},
				{Kind: Identifier, Value: "abc", Row: 2, Col: 1},
				{Kind: CarriageReturn, Value: "\r", Row: 2, Col: 4},
				{Kind: EOF, Value: "", Row: 2, Col: 5},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			sc := NewScanner(tc.input)
			for _, expectedToken := range tc.expected {
				token := sc.NextToken()
				if token.Kind != expectedToken.Kind || token.Value != expectedToken.Value || token.Row != expectedToken.Row || token.Col != expectedToken.Col {
					t.Errorf("expected %#v, got %#v", expectedToken, token)
				}
			}
		})
	}

}
