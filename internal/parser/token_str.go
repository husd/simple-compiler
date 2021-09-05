package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 字符串token
 */
type StringToken struct {
	tk *tokenKind

	lineNum int // 多少行
	linePos int // 位置

	val string //字面量

	pos    int // 开始位置
	endPos int //结束位置

	inx int //符号表里的索引
}

func newStringToken(tk *tokenKind, lineNum int, linePos int,
	val string, pos int, endPos int) *StringToken {

	res := &StringToken{tk, lineNum, linePos,
		val, pos, endPos, -1}
	return res
}

func (st *StringToken) GetTokenKind() *tokenKind {

	return st.tk
}

func (st *StringToken) GetName() *util.Name {

	panic("implement me")
}

func (st *StringToken) GetStringVal() string {

	return st.val
}

func (st *StringToken) GetRadix() int {

	panic("implement me")
}

func (dt *StringToken) DebugToString() string {

	return fmt.Sprintf("stringtoken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (st *StringToken) CheckTokenKind() {
	if st.tk.Tag != TOKEN_TAG_STRING {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_STRING))
	}
}

func (dt *StringToken) Pos() int {

	return dt.pos
}

func (dt *StringToken) EndPos() int {

	return dt.endPos
}

func (dt *StringToken) GetSymbolTableIndex() int {

	return dt.inx
}

func (dt *StringToken) SetSymbolTableIndex(inx int) {

	dt.inx = inx
}
