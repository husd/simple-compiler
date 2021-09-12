package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/lang"
)

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
type JCStatement struct {
	*JCTree
}

func (jc *JCStatement) TreeType() ast_tree2.TreeType {

	panic("implement me")
}

func (jc *JCStatement) StatementTreeV2_() {

	panic("implement me")
}

func NewJCStatement(tree *JCTree) *JCStatement {

	return &JCStatement{tree}
}

func (jc *JCStatement) SetType(javaType *lang.JavaType) *JCStatement {

	jc.JavaType = javaType
	return jc
}

func (jc *JCStatement) SetPos(pos int) *JCStatement {

	jc.Pos = pos
	return jc
}
