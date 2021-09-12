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

	endPosMap map[*jc.JCTree]int
}

func NewSimpleEndPosTable(parser *JavacParser) *SimpleEndPosTable {

	table := SimpleEndPosTable{}

	table.parser = parser
	table.errorEndPos = 0
	table.endPosMap = make(map[*jc.JCTree]int)

	return &table
}

func (table *SimpleEndPosTable) GetEndPos(jcTree *jc.JCTree) int {

	if v, ok := table.endPosMap[jcTree]; ok {
		return v
	}
	return util.POSITION_NOPOS
}

func (table *SimpleEndPosTable) SetEnd(jcTree *jc.JCTree, endPos int) {

	table.endPosMap[jcTree] = endPos
}

func (table *SimpleEndPosTable) ReplaceTree(oldTree *jc.JCTree, newTree *jc.JCTree) int {

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

func (table *SimpleEndPosTable) toP(jcTree *jc.JCTree) *jc.JCExpression {

	table.SetEnd(jcTree, table.parser.token.EndPos())
	return jc.NewJCExpression(jcTree)
}
