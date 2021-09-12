package jc

import (
	"husd.com/v0/util"
)

/**
 * The factory to be used for abstract syntax tree construction.
 * @author hushengdong
 */
type AstTreeMaker struct {
	pos int //在源代码里的位置
}

func NewAstTreeMaker(c *util.Context) *AstTreeMaker {

	treeMaker := AstTreeMaker{0}
	c.Put(util.C_TREE_MAKER, &treeMaker)
	return &treeMaker
}

func InstanceAstTreeMaker(c *util.Context) *AstTreeMaker {

	ok, obj := c.Get(util.C_TREE_MAKER)
	if ok {
		return obj.(*AstTreeMaker)
	}
	return NewAstTreeMaker(c)
}

//指定位置
func (treeMaker *AstTreeMaker) At(pos int) *AstTreeMaker {

	treeMaker.pos = pos
	return treeMaker
}

func (treeMaker *AstTreeMaker) Identify(name *util.Name) *JCIdent {

	ident := NewJCIdent(name)
	ident.pos = treeMaker.pos
	return ident
}

func (treeMaker *AstTreeMaker) Select(selected *JCExpression, selector *util.Name) *JCFieldAccess {

	return NewJCFieldAccess(selected, selector)
}