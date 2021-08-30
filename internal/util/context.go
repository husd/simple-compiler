package util

import "husd.com/v0/parser"

/**
 * 上下文，自定义的上下文，类似ThreadLocal
 */

type ContextKey int

type Context struct {
	Parser parser.Parser // javac_parser
	Tks    *parser.Tokens
}

func NewContext() *Context {

	c := Context{}
	return &c
}
