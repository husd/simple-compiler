package parser

import (
	"fmt"
	"husd.com/v0/common"
	"husd.com/v0/util"
)

/**
 * true false enum void this super byte short char int long float double等
 * 这些都统归为 NamedToken 表示的是定义变量的属性的，可以看到大部分都是基本类型
 */
type NamedToken struct {
	tk      common.TokenKind
	lineNum int // 多少行
	linePos int // 位置
	name    *util.Name

	pos    int // 开始位置
	endPos int //结束位置

	inx int // 符号表里的索引
}

func newNamedToken(tk common.TokenKind, lineNum int, linePos int,
	n *util.Name, pos int, endPos int) *NamedToken {

	res := &NamedToken{tk, lineNum, linePos,
		n, pos, endPos, -1}
	return res
}

func (dt *NamedToken) GetRowNum() int {

	return dt.lineNum
}

func (dt *NamedToken) GetColumnNum() int {

	return dt.linePos
}

func (nt *NamedToken) GetTokenKind() common.TokenKind {

	return nt.tk
}

func (nt *NamedToken) GetName() *util.Name {

	return nt.name
}

func (nt *NamedToken) GetStringVal() string {

	return nt.name.NameStr
}

func (nt *NamedToken) GetRadix() int {

	panic("implement me")
}

func (dt *NamedToken) DebugToString() string {

	return fmt.Sprintf("NamedToken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (nt *NamedToken) CheckTokenKind() {

	if common.GetTokenKindTag(nt.tk) != common.TOKEN_TAG_NAMED {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", common.TOKEN_TAG_NAMED))
	}
}

func (dt *NamedToken) Pos() int {

	return dt.pos
}

func (dt *NamedToken) EndPos() int {

	return dt.endPos
}

func (dt *NamedToken) GetSymbolTableIndex() int {

	return dt.inx
}

func (dt *NamedToken) SetSymbolTableIndex(inx int) {

	dt.inx = inx
}
