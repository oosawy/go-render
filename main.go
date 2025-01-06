package main

import (
	"context"
	"fmt"
	"time"

	x "github.com/oosawy/go-render/gox"
	renderer "github.com/oosawy/go-render/renderer"
)

func message(_ x.Context, props x.Props, _ x.Children) x.Node {
	return x.V(fmt.Sprintf("Hello, %s! The time is %s.", props["name"], time.Now().Format("15:04:05")))
}

func app(_ x.Context, _ x.Props, _ x.Children) x.Node {
	return x.T("div", nil, x.T("h1", nil, x.E(message, x.P{"name": "world"})))
}

func main() {
	ctx := context.Background()
	tree := renderer.Render(ctx, x.E(app, nil))
	fmt.Printf("%+v\n", tree)
}
