package parser

import "github.com/primo-js/template-compiler/internal/scanner"

type Parser struct {
	sc           *scanner.Scanner
	currentToken scanner.Token
}

func (p *Parser) Parse() ([]Node, error) {
	roots := make([]Node, 0)

	for p.currentToken.Kind != scanner.EOF {
		staticText := p.staticText()
		roots = append(roots, staticText)
	}

	return roots, nil
}

func (p *Parser) advance() {
	p.currentToken = p.sc.NextToken()
}

func (p *Parser) staticText() *StaticTextNode {
	staticText := &StaticTextNode{Value: ""}
	for p.currentToken.Kind != scanner.EOF {
		staticText.Value += p.currentToken.Value
		p.advance()
	}
	return staticText
}

func NewParser(sc *scanner.Scanner) *Parser {
	p := &Parser{
		sc:           sc,
		currentToken: scanner.Token{},
	}
	p.advance()
	return p
}
