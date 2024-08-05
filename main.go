package main

import (
	"fmt"
	"github.com/kerimcanbalkan/markdown-converter/parser"
	"github.com/kerimcanbalkan/markdown-converter/renderer"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage markdown-converter <file>")
		return
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read the file %s", err)
		return
	}

	document := parser.Parse(string(content))

	html := renderer.Render(document)

	fmt.Println(html)
}
