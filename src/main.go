package main

import (
	"fmt"

	node "github.com/ErikHedman23/StaticSiteGenerator/nodes"
)

func main() {
	fmt.Println("Hello, world!")

	var newTextNode = node.NewTextNode(0, "Hello, world!")
	fmt.Println(*newTextNode)

	isEqual := newTextNode.Equals(node.NewTextNode(0, "Hello, world!"))
	fmt.Println(isEqual)

	leafNode := node.NewLeafNode("div", "Hello, world!", map[string]string{"class": "test"})
	fmt.Println(leafNode.ToString())

	displayLeafNode, err := leafNode.ToHtml()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(displayLeafNode)
}
