package gox

import "fmt"

type P = Props
type C = Children
type N = Node

// primitive Value node constructor
func V(value any) Node {
	switch value.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return &primNode{value}
	default:
		panic(fmt.Sprintf("x.V: invalid type for primitive Node: %T", value))
	}
}

// Tag element constructor
func T(tag string, props P, children ...Node) Element {
	return &tagElement{tag, props, children}
}

// component Element constructor
func E(comp CompFunc, props P, children ...Node) Element {
	return &compElement{comp, props, children}
}
