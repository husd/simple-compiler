package ast_tree

import "husd.com/v0/lang"

/**
 * statement 语句主要描述的是代码里的控制语句，例如：
 * break
 * continue
 * if () {} else {}
 * for()
 * do {} while()
 * return
 * try {} catch() {} finally {}
 * switch () { case () : }
 *
 * @author hushengdong
 */

type StatementTree interface {
}

type JCStatement struct {
	jcTree *JCTree
}

func NewJCStatement(tree *JCTree) *JCStatement {

	return &JCStatement{jcTree: tree}
}

func (jc *JCStatement) SetType(javaType *lang.JavaType) *JCStatement {

	jc.jcTree.SetJavaType(javaType)
	return jc
}

func (jc *JCStatement) SetPos(pos int) *JCStatement {

	jc.jcTree.SetPos(pos)
	return jc
}
