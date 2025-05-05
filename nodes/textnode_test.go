package nodes

import (
	"testing"
)

func TestTextNodeEquals(t *testing.T) {
	textNode1 := NewTextNode(Normal, "Hello, world!")
	textNode2 := NewTextNode(Normal, "Hello, world!")

	if !textNode1.Equals(textNode2) {
		t.Errorf("Text nodes should be equal")
	}
}

func TestTextNodeNotEquals(t *testing.T) {
	textNode1 := NewTextNode(Normal, "Hello, world!")
	textNode2 := NewTextNode(Bold, "Hello, world!")

	if textNode1.Equals(textNode2) {
		t.Errorf("Text nodes should not be equal")
	}
}
