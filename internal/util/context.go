package util

/**
 * 上下文，自定义的上下文，类似ThreadLocal
 */

type ContextKey int

type Context struct {
}

func NewContext() *Context {

	c := Context{}
	return &c
}
