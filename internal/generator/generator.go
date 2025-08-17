package generator

import (
	"fmt"
	"strings"

	"github.com/primo-js/template-compiler/internal/ast"
	"github.com/primo-js/template-compiler/internal/parser"
	"github.com/primo-js/template-compiler/internal/scanner"
)

type Generator interface {
	Generate() (string, error)
	GenerateStaticText(node *ast.StaticTextNode)
	GenerateInterpolation(node *ast.InterpolationNode)
}

type defaultGenerator struct {
	p                *parser.Parser
	createStatements []string
	mountStatements  []string
	updateStatements []string
}

func (g *defaultGenerator) Generate() (string, error) {

	roots, err := g.p.Parse()
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}
	for _, root := range roots {
		g.generateNode(root)
	}

	var sb strings.Builder
	sb.WriteString("export function template() {")
	if len(g.createStatements) > 0 {
		sb.WriteString("\n")
		for _, stmt := range g.createStatements {
			sb.WriteString(fmt.Sprintf("\t%s\n", stmt))
		}
	}
	if len(g.mountStatements) > 0 {
		sb.WriteString("\treturn {")
		sb.WriteString("\n\t\tmount(container) {\n")
		for _, stmt := range g.mountStatements {
			sb.WriteString(fmt.Sprintf("\t\t\t%s\n", stmt))
		}
		sb.WriteString("\t\t}")
		if len(g.updateStatements) > 0 {
			sb.WriteString(",\n\t\tupdate(changes) {\n")
			for _, stmt := range g.updateStatements {
				sb.WriteString(fmt.Sprintf("\t\t\t%s\n", stmt))
			}
			sb.WriteString("\t\t}")
		}
		sb.WriteString("\n\t}\n")
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func (g *defaultGenerator) generateNode(node ast.Node) {
	switch n := node.(type) {
	case *ast.StaticTextNode:
		g.GenerateStaticText(n)
	case *ast.InterpolationNode:
		g.GenerateInterpolation(n)
	default:
		// Handle other node types as needed
		fmt.Printf("Unhandled node type: %T\n", n)
	}
}

func (g *defaultGenerator) GenerateStaticText(n *ast.StaticTextNode) {
	g.addCreateStatement(fmt.Sprintf("const text = document.createTextNode(\"%s\");", n.Value))
	g.addMountStatement("container.appendChild(text);")
}

func (g *defaultGenerator) GenerateInterpolation(n *ast.InterpolationNode) {
	g.addCreateStatement("const text = document.createTextNode(\"\");")
	g.addMountStatement("container.appendChild(text);")
	g.addUpdateStatement(fmt.Sprintf("text.data = changes[\"%s\"] ?? \"\";", n.Expression.String()))
}

func (g *defaultGenerator) addCreateStatement(stmt string) {
	g.createStatements = append(g.createStatements, stmt)
}

func (g *defaultGenerator) addMountStatement(stmt string) {
	g.mountStatements = append(g.mountStatements, stmt)
}

func (g *defaultGenerator) addUpdateStatement(stmt string) {
	g.updateStatements = append(g.updateStatements, stmt)
}

func NewGenerator(input string) Generator {
	sc := scanner.NewScanner(input)
	p := parser.NewParser(sc)
	g := &defaultGenerator{
		p:                p,
		createStatements: make([]string, 0),
		mountStatements:  make([]string, 0),
	}
	return g
}
