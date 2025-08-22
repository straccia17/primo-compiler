package ast

import (
	"fmt"

	"github.com/primo-js/template-compiler/internal/scanner"
)

type NodeKind int

const (
	StaticText NodeKind = iota
	Interpolation

	Integer
	Identifier
	Binary
)

type Node interface {
	Kind() NodeKind
}

type StaticTextNode struct {
	Value string
}

func (n *StaticTextNode) Kind() NodeKind {
	return StaticText
}

type InterpolationNode struct {
	Expression Expression
}

func (n *InterpolationNode) Kind() NodeKind {
	return Interpolation
}

type Expression interface {
	Node
	String() string
}

type IntegerLiteral struct {
	Value int
}

func (e *IntegerLiteral) Kind() NodeKind {
	return Integer
}

func (e *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", e.Value)
}

type IdentifierExpression struct {
	Value string
}

func (e *IdentifierExpression) Kind() NodeKind {
	return Identifier
}

func (e *IdentifierExpression) String() string {
	return e.Value
}

type BinaryExpression struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (e *BinaryExpression) Kind() NodeKind {
	return Binary
}

func (e *BinaryExpression) String() string {
	return e.Left.String() + " " + e.Operator.Value + " " + e.Right.String()
}
