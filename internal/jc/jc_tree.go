package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/lang"
)

/**
 * JCTree是树的基本操作
 * @author hushengdong
 */
type JCTree struct {
	Pos      int // 源代码里的位置
	JavaType *lang.JavaType

	// abstract method
	AstTreeType func() ast_tree2.TreeType
	getTag      func() JCTreeTag
}

func (jc *JCTree) hasTag(tag JCTreeTag) bool {

	return tag == jc.getTag()
}

func (jc *JCTree) setType(t *lang.JavaType) *JCTree {

	jc.JavaType = t
	return jc
}

func (jc *JCTree) pos() DiagnosticPosition {

	return jc
}

func (jc *JCTree) Cloneable_() {

	panic("implement me")
}

func (jc *JCTree) TreeType() ast_tree2.TreeType {

	return jc.AstTreeType()
}

// DiagnosticPosition
func (jc *JCTree) getTree() *JCTree {
	return jc
}

func (jc *JCTree) getStartPosition() int {
	panic(" getStartPosition implement me")
}

func (jc *JCTree) getPreferredPosition() int {
	return jc.Pos
}

func (jc *JCTree) getEndPosition(endPosTable *EndPosTable) int {
	panic("implement me")
}

// DiagnosticPosition
