package parser

import (
	"husd.com/v0/jc"
	"husd.com/v0/util"
)

/**
 * 一个简单的用map实现的end pos table
 * @author hushengdong
 */
type SimpleEndPosTable struct {
	parser      *JavacParser
	errorEndPos int

	endPosMap map[*jc.AbstractJCTree]int
}

func NewSimpleEndPosTable(parser *JavacParser) *SimpleEndPosTable {

	table := SimpleEndPosTable{}

	table.parser = parser
	table.errorEndPos = 0
	table.endPosMap = make(map[*jc.AbstractJCTree]int)

	return &table
}

func (table *SimpleEndPosTable) GetEndPos(jcTree *jc.AbstractJCTree) int {

	if v, ok := table.endPosMap[jcTree]; ok {
		return v
	}
	return util.POSITION_NOPOS
}

func (table *SimpleEndPosTable) SetEnd(jcTree *jc.AbstractJCTree, endPos int) {

	table.endPosMap[jcTree] = endPos
}

func (table *SimpleEndPosTable) ReplaceTree(oldTree *jc.AbstractJCTree, newTree *jc.AbstractJCTree) int {

	oldPos := table.GetEndPos(oldTree)
	delete(table.endPosMap, oldTree) //从map里移除老的JCTree
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

func (table *SimpleEndPosTable) toP(jcTree *jc.AbstractJCTree) {

	table.SetEnd(jcTree, table.parser.token.EndPos())
}
