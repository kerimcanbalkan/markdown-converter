package renderer

import (
	"fmt"
	"github.com/kerimcanbalkan/markdown-converter/parser"
	"strings"
)

func Render(node *parser.Node) string {
	var sb strings.Builder

	for _, child := range node.Children {
		renderNode(&sb, child)
	}

	return sb.String()
}

func renderNode(sb *strings.Builder, node *parser.Node) {
	switch node.Type {
	case parser.HeadingNode:
		renderHeading(sb, node)
	case parser.ParagraphNode:
		sb.WriteString("<p>")
		sb.WriteString(node.Content)
		sb.WriteString("<p>")
	case parser.BlockquoteNode:
		sb.WriteString("<blockquote>")
		sb.WriteString(node.Content)
		sb.WriteString("</blockquote>")
	case parser.TextNode:
		sb.WriteString(node.Content)
	}
}

func renderHeading(sb *strings.Builder, node *parser.Node) {
	sb.WriteString(fmt.Sprintf("<h%d>", node.Level))
	sb.WriteString(node.Content)
	sb.WriteString(fmt.Sprintf("</h%d>", node.Level))
}
