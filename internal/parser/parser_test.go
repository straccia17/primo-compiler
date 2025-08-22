package parser

import (
	"testing"

	"github.com/straccia17/primo-compiler/internal/ast"
	"github.com/straccia17/primo-compiler/internal/scanner"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []ast.Node
	}{
		{
			"empty_template",
			"",
			[]ast.Node{},
		},
		{
			"static_text",
			"Hello, world!",
			[]ast.Node{
				&ast.StaticTextNode{Value: "Hello, world!"},
			},
		},
		{
			"interpolation",
			"{ name }",
			[]ast.Node{
				&ast.InterpolationNode{
					Expression: &ast.IdentifierExpression{
						Value: "name",
					},
				},
			},
		},
		{
			"addition and multiplication",
			"{ 1 + 2 * 3 }",
			[]ast.Node{
				&ast.InterpolationNode{
					Expression: &ast.BinaryExpression{
						Left:     &ast.IntegerLiteral{Value: 1},
						Operator: simpleToken(scanner.Plus, "+"),
						Right: &ast.BinaryExpression{
							Left:     &ast.IntegerLiteral{Value: 2},
							Operator: simpleToken(scanner.Star, "*"),
							Right:    &ast.IntegerLiteral{Value: 3},
						},
					},
				},
			},
		},
		{
			"subtraction",
			"{ 1 - 2 }",
			[]ast.Node{
				&ast.InterpolationNode{
					Expression: &ast.BinaryExpression{
						Left:     &ast.IntegerLiteral{Value: 1},
						Operator: simpleToken(scanner.Minus, "-"),
						Right:    &ast.IntegerLiteral{Value: 2},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := scanner.NewScanner(tt.input)
			p := NewParser(sc)
			result, err := p.Parse()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d nodes, got %d", len(tt.expected), len(result))
			}

			for i, expected := range tt.expected {
				if result[i].Kind() != expected.Kind() {
					t.Errorf("expected node kind %v, got %v", expected.Kind(), result[i].Kind())
				}

				switch n := expected.(type) {
				case *ast.StaticTextNode:
					if staticTextNode, ok := result[i].(*ast.StaticTextNode); ok {
						if staticTextNode.Value != n.Value {
							t.Errorf("expected static text '%s', got '%s'", n.Value, staticTextNode.Value)
						}
					}
				case *ast.InterpolationNode:
					if interpolationNode, ok := result[i].(*ast.InterpolationNode); ok {
						if interpolationNode.Expression.String() != n.Expression.String() {
							t.Errorf("expected interpolation expression '%s', got '%s'", n.Expression.String(), interpolationNode.Expression.String())
						}
					}
				}
			}
		})
	}
}

func simpleToken(t scanner.TokenKind, v string) scanner.Token {
	return scanner.Token{Kind: t, Value: v}
}
