package util

/**
 * 上下文，自定义的上下文，类似ThreadLocal
 * 目前还没用上，之后要用上，要解决一个类型存储的问题。 TODO husd
 * @author hushengdong
 */
type ContextKey int

type Context struct {
}

func NewContext() *Context {

	c := Context{}
	return &c
}
