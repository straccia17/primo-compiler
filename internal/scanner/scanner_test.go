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
				{kind: EOF, value: "", row: 1, col: 1},
			},
		},
		{
			"operators",
			"+-*/%<>=!&|^?",
			[]Token{
				{kind: Plus, value: "+", row: 1, col: 1},
				{kind: Minus, value: "-", row: 1, col: 2},
				{kind: Star, value: "*", row: 1, col: 3},
				{kind: Slash, value: "/", row: 1, col: 4},
				{kind: Percent, value: "%", row: 1, col: 5},
				{kind: LessThan, value: "<", row: 1, col: 6},
				{kind: GreaterThan, value: ">", row: 1, col: 7},
				{kind: Equal, value: "=", row: 1, col: 8},
				{kind: Not, value: "!", row: 1, col: 9},
				{kind: And, value: "&", row: 1, col: 10},
				{kind: Or, value: "|", row: 1, col: 11},
				{kind: Caret, value: "^", row: 1, col: 12},
				{kind: Question, value: "?", row: 1, col: 13},
				{kind: EOF, value: "", row: 1, col: 14},
			},
		},
		{
			"numbers",
			"123 456.789",
			[]Token{
				{kind: Integer, value: "123", row: 1, col: 1},
				{kind: Whitespace, value: " ", row: 1, col: 4},
				{kind: Float, value: "456.789", row: 1, col: 5},
				{kind: EOF, value: "", row: 1, col: 13},
			},
		},
		{
			"identifiers",
			"var1 var_2 var3",
			[]Token{
				{kind: Identifier, value: "var1", row: 1, col: 1},
				{kind: Whitespace, value: " ", row: 1, col: 5},
				{kind: Identifier, value: "var_2", row: 1, col: 6},
				{kind: Whitespace, value: " ", row: 1, col: 12},
				{kind: Identifier, value: "var3", row: 1, col: 13},
				{kind: EOF, value: "", row: 1, col: 17},
			},
		},
		{
			"parentheses",
			"()[]{}",
			[]Token{
				{kind: LeftParen, value: "(", row: 1, col: 1},
				{kind: RightParen, value: ")", row: 1, col: 2},
				{kind: LeftBracket, value: "[", row: 1, col: 3},
				{kind: RightBracket, value: "]", row: 1, col: 4},
				{kind: LeftBrace, value: "{", row: 1, col: 5},
				{kind: RightBrace, value: "}", row: 1, col: 6},
				{kind: EOF, value: "", row: 1, col: 7},
			},
		},
		{
			"punctuation",
			",.;:",
			[]Token{
				{kind: Comma, value: ",", row: 1, col: 1},
				{kind: Dot, value: ".", row: 1, col: 2},
				{kind: Semicolon, value: ";", row: 1, col: 3},
				{kind: Colon, value: ":", row: 1, col: 4},
				{kind: EOF, value: "", row: 1, col: 5},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			sc := NewScanner(tc.input)
			for _, expectedToken := range tc.expected {
				token := sc.NextToken()
				if token.kind != expectedToken.kind || token.value != expectedToken.value {
					t.Errorf("expected %#v, got %#v", expectedToken, token)
				}
			}
		})
	}

}
