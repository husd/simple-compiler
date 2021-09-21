package jc

import (
	"husd.com/v0/ast_tree2"
)

/**
 * expression主要描述了代码里的表达式，例如：
 * 三元表达式 bool ? () : ()
 * 赋值语句： () = ()
 * New操作 () () = new ()
 *
 * expression是代码解析里比较基础的类型
 * @author hushengdong
 */
type AbstractJCExpression struct {
	*AbstractJCTree
}

func (jc *AbstractJCExpression) TreeType() *ast_tree2.TreeType {

	return jc.getTreeType()
}

func (jc *AbstractJCExpression) ExpressionTreeV2_() {

	//panic("implement me")
}

func NewJCExpression() *AbstractJCExpression {

	res := &AbstractJCExpression{}
	res.AbstractJCTree = NewJCTree()
	return res
}
