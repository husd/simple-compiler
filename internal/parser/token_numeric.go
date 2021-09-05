package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 数字类的token
 */
type NumericToken struct {
	tk      *tokenKind
	lineNum int // 多少行
	linePos int // 位置

	val   string
	radix int

	pos    int // 开始位置
	endPos int //结束位置

	inx int //符号表里的索引
}

func newNumericToken(tk *tokenKind, lineNum int, linePos int,
	val string, radix int, pos int, endPos int) *NumericToken {

	res := &NumericToken{tk, lineNum, linePos,
		val, radix, pos, endPos, -1}
	return res
}

func (nt *NumericToken) GetTokenKind() *tokenKind {
	return nt.tk
}

func (nt *NumericToken) GetName() *util.Name {

	n := util.Name{NameStr: nt.val, Index: 0}
	return &n
}

func (nt *NumericToken) GetStringVal() string {

	return nt.val
}

func (nt *NumericToken) GetRadix() int {
	return nt.radix
}

func (dt *NumericToken) DebugToString() string {

	return fmt.Sprintf("NumericToken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (nt *NumericToken) CheckTokenKind() {
	if nt.tk.Tag != TOKEN_TAG_NUMERIC {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_NUMERIC))
	}
}

func (dt *NumericToken) Pos() int {

	return dt.pos
}

func (dt *NumericToken) EndPos() int {

	return dt.endPos
}

func (dt *NumericToken) GetSymbolTableIndex() int {

	return dt.inx
}

func (dt *NumericToken) SetSymbolTableIndex(inx int) {

	dt.inx = inx
}
