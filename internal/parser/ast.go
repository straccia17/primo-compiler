package parser

type NodeKind int

const (
	StaticText NodeKind = iota
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
