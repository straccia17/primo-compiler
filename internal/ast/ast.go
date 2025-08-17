package ast

type NodeKind int

const (
	StaticText NodeKind = iota
	Interpolation

	Identifier
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

type IdentifierExpression struct {
	Value string
}

func (e *IdentifierExpression) Kind() NodeKind {
	return Identifier
}

func (e *IdentifierExpression) String() string {
	return e.Value
}
