package jc

import (
	"husd.com/v0/code"
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

func (treeMaker *AstTreeMaker) Identify(name *util.Name) *AbstractJCExpression {

	ident := NewJCIdent(name)
	ident.pos = treeMaker.pos
	return ident.AbstractJCExpression
}

func (treeMaker *AstTreeMaker) Select(selected *AbstractJCExpression, selector *util.Name) *JCFieldAccess {

	tree := NewJCFieldAccess(selected, selector, nil)
	tree.Pos = treeMaker.pos
	return tree
}

func (treeMaker *AstTreeMaker) Literal(tag *code.TypeTag, n int) *JCLiteral {

	res := NewJCLiteral(tag, n)
	res.Pos = treeMaker.pos
	return res
}

func (treeMaker *AstTreeMaker) TypeIdent(tag *code.TypeTag) *JCPrimitiveTypeTree {

	pt := NewJCPrimitiveTypeTree(tag)
	pt.Pos = treeMaker.pos
	return pt
}

func (treeMaker *AstTreeMaker) Indexed(t *AbstractJCExpression, t1 *AbstractJCExpression) *JCArrayAccess {

	tree := NewJCArrayAccess(t, t1)
	tree.Pos = treeMaker.pos
	return tree
}

func (treeMaker *AstTreeMaker) Assign(left *AbstractJCExpression, right *AbstractJCExpression) *JCAssign {

	tree := NewJCAssign(left, right)
	tree.Pos = treeMaker.pos
	return tree
}

func (treeMaker *AstTreeMaker) Assignop(tag JCTreeTag, left *AbstractJCExpression, right *AbstractJCExpression) *JCAssignOp {

	tree := NewJCAssignOp(tag, left, right, nil)
	tree.Pos = treeMaker.pos
	return tree
}

func (treeMaker *AstTreeMaker) Conditional(condition *AbstractJCExpression, truePart *AbstractJCExpression, falsePart *AbstractJCExpression) *JCConditional {

	tree := NewJCConditional(condition, truePart, falsePart)
	tree.Pos = treeMaker.pos
	return tree
}
