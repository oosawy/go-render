package gox

import "context"

type Context struct {
	context.Context
}

func WithContext(ctx context.Context) Context {
	return Context{ctx}
}
