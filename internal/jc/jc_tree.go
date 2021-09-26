package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/lang"
)

/**
 * JCTree是树的基本操作
 * @author hushengdong
 */
type AbstractJCTree struct {
	Pos      int // 源代码里的位置
	JavaType *lang.JavaType

	// abstract method
	getTreeType func() *ast_tree2.TreeType
	getTag      func() JCTreeTag
	toString    func() string
}

func NewJCTree() *AbstractJCTree {

	tree := &AbstractJCTree{}

	tree.toString = func() string {
		return tree.getTreeType().Name
	}

	return tree
}

func (jc *AbstractJCTree) hasTag(tag JCTreeTag) bool {

	return tag == jc.getTag()
}

func (jc *AbstractJCTree) setType(t *lang.JavaType) *AbstractJCTree {

	jc.JavaType = t
	return jc
}

func (jc *AbstractJCTree) pos() DiagnosticPosition {

	return jc
}

func (jc *AbstractJCTree) Cloneable_() {
}

func (jc *AbstractJCTree) TreeType() *ast_tree2.TreeType {

	return jc.getTreeType()
}

// DiagnosticPosition
func (jc *AbstractJCTree) getTree() *AbstractJCTree {
	return jc
}

func (jc *AbstractJCTree) getStartPosition() int {
	panic(" getStartPosition implement me")
}

func (jc *AbstractJCTree) getPreferredPosition() int {
	return jc.Pos
}

func (jc *AbstractJCTree) getEndPosition(endPosTable *EndPosTable) int {
	panic("implement me")
}

// DiagnosticPosition
