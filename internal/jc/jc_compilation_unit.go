package jc

import (
	"husd.com/v0/lang"
)

/**
 * 源代码里的所有东西在这里都有 这个结构是一个大而全的结构
 * @author hushengdong
 */
type JCCompilationUnit struct {
	*JCTree
}

func NewJCCompilationUnit(tree *JCTree) *JCCompilationUnit {

	return &JCCompilationUnit{tree}
}

func (jc *JCCompilationUnit) SetType(javaType *lang.JavaType) *JCCompilationUnit {

	jc.JavaType = javaType
	return jc
}

func (jc *JCCompilationUnit) SetPos(pos int) *JCCompilationUnit {

	jc.Pos = pos
	return jc
}
