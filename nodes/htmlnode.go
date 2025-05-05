package nodes

import (
	"fmt"
	"strings"
)

type HtmlNode struct {
	Tag      *string
	Value    *string
	Children []*HtmlNode
	Props    map[string]string
}

func NewHtmlNode(tag, value *string, children []*HtmlNode, props map[string]string) *HtmlNode {
	return &HtmlNode{
		Tag:      tag,
		Value:    value,
		Children: children,
		Props:    props,
	}
}

func (html *HtmlNode) ToString() string {
	return fmt.Sprintf("<%s%s>%s</%s>", *html.Tag, html.PropsToHtml(), *html.Value, *html.Tag)
}

func (node *HtmlNode) ToHtml() (string, error) {
	return "", fmt.Errorf("method not implemented")
}

func (node *HtmlNode) PropsToHtml() string {
	var sb strings.Builder
	for key, value := range node.Props {
		sb.WriteString(fmt.Sprintf(`%s="%s"`, key, value))
	}
	return sb.String()
}
