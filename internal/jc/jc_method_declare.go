package jc

import (
	"husd.com/v0/lang"
)

/**
 * 这个主要是方法声明
 * @author hushengdong
 */
type JCMethodDecl struct {
	*JCTree
}

func NewJCMethodDecl(tree *JCTree) *JCMethodDecl {

	return &JCMethodDecl{tree}
}

func (jc *JCMethodDecl) SetType(javaType *lang.JavaType) *JCMethodDecl {

	jc.JavaType = javaType
	return jc
}

func (jc *JCMethodDecl) SetPos(pos int) *JCMethodDecl {

	jc.Pos = pos
	return jc
}
