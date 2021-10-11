package parser

import (
	"husd.com/v0/ast"
	"husd.com/v0/util"
)

/**
 * 一个简单的用map实现的end pos table
 * @author hushengdong
 */
type SimpleEndPosTable struct {
	parser      *JavacParser
	errorEndPos int

	endPosMap map[*ast.TreeNode]int
}

func NewSimpleEndPosTable(parser *JavacParser) *SimpleEndPosTable {

	table := SimpleEndPosTable{}

	table.parser = parser
	table.errorEndPos = 0
	table.endPosMap = make(map[*ast.TreeNode]int)

	return &table
}

func (table *SimpleEndPosTable) GetEndPos(jcTree *ast.TreeNode) int {

	if v, ok := table.endPosMap[jcTree]; ok {
		return v
	}
	return util.POSITION_NOPOS
}

func (table *SimpleEndPosTable) SetEnd(jcTree *ast.TreeNode, endPos int) {

	table.endPosMap[jcTree] = endPos
}

func (table *SimpleEndPosTable) ReplaceTree(oldTree *ast.TreeNode, newTree *ast.TreeNode) int {

	oldPos := table.GetEndPos(oldTree)
	delete(table.endPosMap, oldTree) // 从map里移除老的JCTree
	if oldPos != util.POSITION_NOPOS {
		table.SetEnd(newTree, oldPos)
	}
	return oldPos
}

func (table *SimpleEndPosTable) SetErrorPos(pos int) {

	if pos > table.errorEndPos {
		table.errorEndPos = pos
	}
}

func (table *SimpleEndPosTable) toP(jcTree *ast.TreeNode) {

	table.SetEnd(jcTree, table.parser.token.EndPos())
}
