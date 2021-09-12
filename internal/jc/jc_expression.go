package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/lang"
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
type JCExpression struct {
	*JCTree
}

func (jc *JCExpression) TreeType() ast_tree2.TreeType {

	panic("implement me")
}

func (jc *JCExpression) ExpressionTreeV2_() {
	panic("implement me")
}

func NewJCExpression(tree *JCTree) *JCExpression {

	return &JCExpression{tree}
}

func (jc *JCExpression) SetType(javaType *lang.JavaType) *JCExpression {

	jc.JavaType = javaType
	return jc
}

func (jc *JCExpression) SetPos(pos int) *JCExpression {

	jc.Pos = pos
	return jc
}

func (jc *JCExpression) GetJCTree() *JCTree {

	return jc.getTree()
}
