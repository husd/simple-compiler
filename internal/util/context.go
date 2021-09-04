package util

/**
 * 上下文，自定义的上下文，类似ThreadLocal
 */

type ContextKey int

type Context struct {
	apps map[string]interface{}
}

const (
	C_TREE_MAKER string = "AstTreeMaker"
	C_NAMES      string = "Names"
	C_LEXER      string = "ScannerLexer"
)

func NewContext() *Context {

	c := Context{}

	c.apps = make(map[string]interface{})
	return &c
}

func (c *Context) Get(name string) (bool, interface{}) {

	if v, ok := c.apps[name]; ok {
		return true, v
	}
	return false, nil
}

func (c *Context) Put(name string, obj interface{}) {

	c.apps[name] = obj
}
