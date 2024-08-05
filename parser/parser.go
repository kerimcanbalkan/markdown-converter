package parser

type NodeType int

const (
	TextNode NodeType = iota
	HeadingNode
	ParagraphNode
	BlockquoteNode
)

type Node struct {
	Type     NodeType
	Level    int
	Content  string
	Children []*Node
}

func Parse(input string) *Node {
	return &Node{
		Type:     TextNode,
		Content:  input,
		Children: parseBlocks(input),
	}
}
