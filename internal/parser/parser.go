package parser

import (
	"errors"
	"strconv"

	"github.com/primo-js/template-compiler/internal/ast"
	"github.com/primo-js/template-compiler/internal/scanner"
)

type Parser struct {
	sc           *scanner.Scanner
	tokens       []scanner.Token
	currentToken scanner.Token
}

func (p *Parser) Parse() ([]ast.Node, error) {
	roots := make([]ast.Node, 0)

	var root ast.Node
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
	p.tokens = append(p.tokens, p.currentToken)
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[len(p.tokens)-2]
}

func (p *Parser) match(kinds ...scanner.TokenKind) bool {
	for _, kind := range kinds {
		if p.currentToken.Kind == kind {
			p.advance()
			return true
		}
	}
	return false
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

func (p *Parser) staticText() *ast.StaticTextNode {
	staticText := &ast.StaticTextNode{Value: ""}
	for p.currentToken.Kind != scanner.EOF {
		staticText.Value += p.currentToken.Value
		p.advance()
	}
	return staticText
}

func (p *Parser) interpolation() (*ast.InterpolationNode, error) {
	p.advance() // Skip the left bracket
	p.skipWhitespace()
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	p.skipWhitespace()                  // Move to the next token after the identifier
	err = p.consume(scanner.RightBrace) // Expect a right bracket
	if err != nil {
		return nil, err
	}
	return &ast.InterpolationNode{Expression: expr}, nil
}

func (p *Parser) expression() (ast.Expression, error) {
	left, err := p.factor()
	if err != nil {
		return nil, err
	}
	p.skipWhitespace()
	if p.match(scanner.Plus, scanner.Minus) {
		operator := p.previous()
		p.skipWhitespace()
		right, err := p.expression()
		if err != nil {
			return nil, err
		}
		return &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}, nil
	}
	return left, nil
}

func (p *Parser) factor() (ast.Expression, error) {
	left, err := p.primitive()
	if err != nil {
		return nil, err
	}
	p.skipWhitespace()
	if p.match(scanner.Star, scanner.Slash, scanner.Percent) {
		operator := p.previous()
		p.skipWhitespace()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		return &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}, nil
	}
	return left, nil
}

func (p *Parser) primitive() (ast.Expression, error) {
	if p.currentToken.Kind == scanner.Integer {
		value, err := strconv.Atoi(p.currentToken.Value)
		if err != nil {
			return nil, err
		}
		literal := &ast.IntegerLiteral{Value: value}
		p.advance()
		return literal, nil
	} else if p.currentToken.Kind == scanner.Identifier {
		ident := &ast.IdentifierExpression{Value: p.currentToken.Value}
		p.advance()
		return ident, nil
	}
	return nil, errors.New("unexpected token: " + p.currentToken.Value)
}

func NewParser(sc *scanner.Scanner) *Parser {
	p := &Parser{
		sc:           sc,
		currentToken: scanner.Token{},
		tokens:       make([]scanner.Token, 0),
	}
	p.advance()
	return p
}
