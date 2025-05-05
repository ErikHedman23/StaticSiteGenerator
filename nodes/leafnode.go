package nodes

import (
	"fmt"
	"strings"
)

type LeafNode struct {
	HtmlNode
}

func NewLeafNode(tag string, value string, props map[string]string) *LeafNode {
	return &LeafNode{
		HtmlNode: HtmlNode{
			Tag:   &tag,
			Value: &value,
			Props: props,
		},
	}
}

func (node *LeafNode) ToString() string {
	return fmt.Sprintf("<%s %s>%s</%s>", *node.Tag, node.PropsToHtml(), *node.Value, *node.Tag)
}

func (node *LeafNode) ToHtml() (string, error) {
	if node.Tag == nil {
		return fmt.Sprintf(*node.Value), nil
	}

	if node.Value == nil {
		return "", fmt.Errorf("Value cannot be nil")
	}

	return fmt.Sprintf("<%s %s>%s</%s>", *node.Tag, node.PropsToHtml(), *node.Value, *node.Tag), nil
}

func (node *LeafNode) PropsToHtml() string {
	var sb strings.Builder
	for key, value := range node.Props {
		sb.WriteString(fmt.Sprintf(`%s="%s"`, key, value))
	}
	return sb.String()
}
