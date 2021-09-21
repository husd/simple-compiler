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
type AbstractJCStatement struct {
	*AbstractJCTree
}

func (jc *AbstractJCStatement) TreeType() ast_tree2.TreeType {

	panic("implement me")
}

func (jc *AbstractJCStatement) StatementTreeV2_() {

	//panic("implement me")
}

func NewJCStatement(tree *AbstractJCTree) *AbstractJCStatement {

	return &AbstractJCStatement{tree}
}

func (jc *AbstractJCStatement) SetType(javaType *lang.JavaType) *AbstractJCStatement {

	jc.JavaType = javaType
	return jc
}

func (jc *AbstractJCStatement) SetPos(pos int) *AbstractJCStatement {

	jc.Pos = pos
	return jc
}
