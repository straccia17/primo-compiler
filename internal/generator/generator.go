package generator

import (
    "fmt"
    "io"
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

type importStatements struct {
    imports map[string][]string
}

func (i *importStatements) AddImport(module, name string) {
    if i.imports == nil {
        i.imports = make(map[string][]string)
    }
    if _, exists := i.imports[module]; !exists {
        i.imports[module] = []string{}
    }
    i.imports[module] = append(i.imports[module], name)
}

func (i *importStatements) Generate(w io.Writer) error {
    if len(i.imports) == 0 {
        return nil
    }

    for module, names := range i.imports {
        if len(names) == 0 {
            continue
        }
        importLine := fmt.Sprintf("import { %s } from \"%s\";\n", strings.Join(names, ", "), module)
        _, err := fmt.Fprintln(w, importLine)
        if err != nil {
            return err
        }

    }
    return nil
}

type defaultGenerator struct {
    p                *parser.Parser
    createStatements []string
    mountStatements  []string
    imports          importStatements
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

    err = g.imports.Generate(&sb)
    if err != nil {
        return "", err
    }

    sb.WriteString("export function template(ctx) {")
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
    g.imports.AddImport("core/signal", "createEffect")
    g.addCreateStatement("const text = document.createTextNode(\"\");")
    g.addCreateStatement(fmt.Sprintf(`createEffect(() => text.data = ctx["%s"]());`, n.Expression.String()))
    g.addMountStatement("container.appendChild(text);")
}

func (g *defaultGenerator) addCreateStatement(stmt string) {
    g.createStatements = append(g.createStatements, stmt)
}

func (g *defaultGenerator) addMountStatement(stmt string) {
    g.mountStatements = append(g.mountStatements, stmt)
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
