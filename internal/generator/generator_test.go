package generator

import (
	"os"
	"path"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		title string
		input string
	}{
		{
			"empty_template",
			"",
		},
		{
			"static_text",
			"Hello, World!",
		},
		{
			"interpolation",
			"{ name }",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// Initialize the generator with the input
			gen := NewGenerator(test.input)

			// Generate the output
			output, err := gen.Generate()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check if the output is as expected
			content, err := os.ReadFile(path.Join("..", "..", "test", "generator", test.title+".js"))
			if err != nil {
				t.Fatalf("failed to read expected output file: %v", err)
			}

			expected := string(content)

			if output != expected {
				t.Errorf("expected output:\n%s\nbut got:\n%s", expected, output)
			}
		})
	}
}
