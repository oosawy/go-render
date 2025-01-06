package html_renderer

import (
	"context"
	"fmt"

	"github.com/oosawy/go-render/gox"
)

func Render(ctx context.Context, elm gox.Element) string {
	tree := gox.RenderTree(ctx, elm)
	return renderNode(ctx, tree)
}

func renderNode(ctx context.Context, node gox.TreeNode) string {
	switch node := node.(type) {
	case gox.TreeElement:
		return renderElement(ctx, node)
	case gox.TreeTextNode:
		return node.Text
	default:
		panic("invalid type")
	}
}

func renderElement(ctx context.Context, elm gox.TreeElement) string {
	children := renderChildren(ctx, elm.Children)
	return fmt.Sprintf("<%s>%s</%s>", elm.Type, children, elm.Type)
}

func renderChildren(ctx context.Context, children []gox.TreeNode) string {
	var renderedChildren string
	for _, child := range children {
		renderedChildren += renderNode(ctx, child)
	}
	return renderedChildren
}
