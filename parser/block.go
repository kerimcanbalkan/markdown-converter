package parser

import (
	"bufio"
	"strings"
)

func parseBlocks(input string) []*Node {
	var nodes []*Node
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		if node := parseHeading(line); node != nil {
			nodes = append(nodes, node)
		} else if node := parseBlockquote(line); node != nil {
			nodes = append(nodes, node)
		} else if len(line) > 0 {
			node := parseParagraph(line)
			nodes = append(nodes, node)
		} else {

		}
	}

	return nodes
}

func parseHeading(line string) *Node {
	if strings.HasPrefix(line, "#") {
		level := 1
		for level < len(line) && line[level] == '#' {
			level++
		}
		if level > 6 {
			return nil
		}

		return &Node{
			Type:    HeadingNode,
			Level:   level,
			Content: strings.TrimSpace(line[level:]),
			Children: []*Node{
				{
					Type:    TextNode,
					Content: strings.TrimSpace(line[level:]),
				},
			},
		}
	}
	return nil
}

func parseParagraph(line string) *Node {
	return &Node{
		Type:    ParagraphNode,
		Content: strings.TrimSpace(line),
		Children: []*Node{
			{
				Type:    TextNode,
				Content: strings.TrimSpace(line),
			},
		},
	}
}

func parseBlockquote(line string) *Node {
	if strings.HasPrefix(line, ">") {
		return &Node{
			Type:    BlockquoteNode,
			Content: strings.TrimSpace(line[1:]),
			Children: []*Node{
				{
					Type:    TextNode,
					Content: strings.TrimSpace(line[1:]),
				},
			},
		}
	}
	return nil
}
