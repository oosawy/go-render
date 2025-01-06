package gox

type Node interface {
	isNode()
	render(Context) TreeNode
}

type primNode struct {
	// value must be a string, numbers or bool
	value any
}

func (*primNode) isNode() {}

type Element interface {
	Node
	isElement()
}

type Props = map[string]any
type Children []Node

type tagElement struct {
	tag      string
	props    Props
	children Children
}

func (*tagElement) isNode()    {}
func (*tagElement) isElement() {}

type CompFunc func(Context, Props, Children) Node
type compElement struct {
	comp     CompFunc
	props    Props
	children Children
}

func (*compElement) isNode()    {}
func (*compElement) isElement() {}
