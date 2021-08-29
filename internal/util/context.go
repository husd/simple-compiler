package util

/**
 * 上下文，自定义的上下文，类似ThreadLocal
 */

type ContextKey int

const (
	context_key_java_compiler ContextKey = 1
)

type Context struct {
	m map[ContextKey]interface{}
}

func NewContext() *Context {

	c := Context{}
	c.m = make(map[ContextKey]interface{})

	return &c
}

/**
 *
 * fac 是工厂函数，用来返回实际需要的对象
 */
func (c *Context) get(key ContextKey, fac func() interface{}) interface{} {

	if o, ok := c.m[key]; ok {
		return o
	}
	o := fac()
	return o
}
