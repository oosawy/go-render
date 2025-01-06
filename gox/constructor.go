package gox

import (
	"context"
	"fmt"
)

const NodeText = "gox$text"

type TreeNode interface {
	isTreeNode()
	isTreeChild()
}

type TreeElement struct {
	Type     string
	Props    map[string]interface{}
	Children []TreeChild
}

func (TreeElement) isTreeNode()  {}
func (TreeElement) isTreeChild() {}

type TreeChild interface {
	isTreeChild()
}

type TreeTextNode struct {
	Text string
}

func (TreeTextNode) isTreeNode()  {}
func (TreeTextNode) isTreeChild() {}

func Render(c context.Context, elm Element) TreeNode {
	ctx := WithContext(c)

	return elm.render(ctx)
}

func (n primNode) render(ctx Context) TreeNode {
	switch n.value.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return TreeTextNode{Text: fmt.Sprintf("%v", n.value)}
	case bool:
		return TreeTextNode{ /* text: "" */ }
	default:
		panic("invalid type")
	}
}

func (c Children) render(ctx Context) []TreeChild {
	children := make([]TreeChild, len(c))
	for i, child := range c {
		children[i] = child.render(ctx)
	}
	return children
}

func (e tagElement) render(ctx Context) TreeNode {
	children := e.children.render(ctx)
	return TreeElement{
		Type:     e.tag,
		Props:    e.props,
		Children: children,
	}
}

func (e compElement) render(ctx Context) TreeNode {
	node := e.comp(ctx, e.props, e.children)
	return node.render(ctx)
}
