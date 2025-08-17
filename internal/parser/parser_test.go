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
		})
	}
}
