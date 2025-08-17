package parser

import (
	"testing"

	"github.com/primo-js/template-compiler/internal/scanner"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Node
	}{
		{
			"empty_template",
			"",
			[]Node{},
		},
		{
			"static_text",
			"Hello, world!",
			[]Node{
				&StaticTextNode{Value: "Hello, world!"},
			},
		},
		{
			"interpolation",
			"{ name }",
			[]Node{
				&InterpolationNode{
					Expression: &IdentifierExpression{
						Value: "name",
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
				case *StaticTextNode:
					if staticTextNode, ok := result[i].(*StaticTextNode); ok {
						if staticTextNode.Value != n.Value {
							t.Errorf("expected static text '%s', got '%s'", n.Value, staticTextNode.Value)
						}
					}
				case *InterpolationNode:
					if interpolationNode, ok := result[i].(*InterpolationNode); ok {
						if interpolationNode.Expression.String() != n.Expression.String() {
							t.Errorf("expected interpolation expression '%s', got '%s'", n.Expression.String(), interpolationNode.Expression.String())
						}
					}
				}
			}
		})
	}
}
