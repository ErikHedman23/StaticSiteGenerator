package nodes

type NodeType int

const (
	Normal NodeType = iota
	Bold
	Italic
	Code
	Link
	Image
)

func (nodeType NodeType) NodeToString() string {
	return [...]string{
		"Normal",
		"Bold",
		"Italic",
		"Code",
		"Link",
		"Image",
	}[nodeType]
}

type TextNode struct {
	Type     NodeType
	Content  string
	Children []*TextNode
}

func NewTextNode(nodeType NodeType, content string) *TextNode {
	return &TextNode{
		Type:     nodeType,
		Content:  content,
		Children: make([]*TextNode, 0),
	}
}

func (n *TextNode) Equals(other *TextNode) bool {
	if n == other {
		return true
	}

	if other == nil {
		return false
	}

	if len(n.Children) != len(other.Children) {
		return false
	}

	for i, child := range n.Children {
		if !child.Equals(other.Children[i]) {
			return false
		}
	}

	return true
}
