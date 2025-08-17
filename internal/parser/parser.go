package parser

import (
	"errors"

	"github.com/primo-js/template-compiler/internal/scanner"
)

type Parser struct {
	sc           *scanner.Scanner
	currentToken scanner.Token
}

func (p *Parser) Parse() ([]Node, error) {
	roots := make([]Node, 0)

	var root Node
	var err error
	for p.currentToken.Kind != scanner.EOF {
		switch p.currentToken.Kind {
		case scanner.LeftBrace:
			root, err = p.interpolation()
		default:
			root = p.staticText()
		}
		if err != nil {
			return nil, err
		}
		roots = append(roots, root)
	}

	return roots, nil
}

func (p *Parser) advance() {
	p.currentToken = p.sc.NextToken()
}

func (p *Parser) skipWhitespace() {
	for p.currentToken.Kind == scanner.Whitespace ||
		p.currentToken.Kind == scanner.Newline ||
		p.currentToken.Kind == scanner.Tab ||
		p.currentToken.Kind == scanner.CarriageReturn {
		p.advance()
	}
}

func (p *Parser) consume(kind scanner.TokenKind) error {
	if p.currentToken.Kind != kind {
		return errors.New("unexpected token: " + p.currentToken.Value + ", expected: " + kind.String())
	}
	p.advance()
	return nil
}

func (p *Parser) staticText() *StaticTextNode {
	staticText := &StaticTextNode{Value: ""}
	for p.currentToken.Kind != scanner.EOF {
		staticText.Value += p.currentToken.Value
		p.advance()
	}
	return staticText
}

func (p *Parser) interpolation() (*InterpolationNode, error) {
	p.advance() // Skip the left bracket
	p.skipWhitespace()
	expr := &IdentifierExpression{Value: p.currentToken.Value}
	p.advance()
	p.skipWhitespace()                   // Move to the next token after the identifier
	err := p.consume(scanner.RightBrace) // Expect a right bracket
	if err != nil {
		return nil, err
	}
	return &InterpolationNode{Expression: expr}, nil
}

func NewParser(sc *scanner.Scanner) *Parser {
	p := &Parser{
		sc:           sc,
		currentToken: scanner.Token{},
	}
	p.advance()
	return p
}
