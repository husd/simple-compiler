package ast_tree

import "husd.com/v0/lang"

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
	jcTree *JCTree
}

func NewJCExpression(tree *JCTree) *JCExpression {

	return &JCExpression{jcTree: tree}
}

func (jc *JCExpression) SetType(javaType *lang.JavaType) *JCExpression {

	jc.jcTree.SetJavaType(javaType)
	return jc
}

func (jc *JCExpression) SetPos(pos int) *JCExpression {

	jc.jcTree.SetPos(pos)
	return jc
}

func (jc *JCExpression) GetJCTree() *JCTree {

	return jc.jcTree
}
