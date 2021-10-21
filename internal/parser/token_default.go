package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 这个是java的关键字和运算符号之类
 * @author hushengdong
 */
type DefaultToken struct {
	tk TokenKind

	lineNum int // 多少行
	linePos int // 位置

	pos    int // 开始位置
	endPos int // 结束位置

	inx int // 符号表里的索引
}

func newDefaultToken(tk TokenKind, lineNum int, linePos int, pos int, endPos int) *DefaultToken {

	res := &DefaultToken{tk, lineNum, linePos, pos, endPos, -1}
	return res
}

func (dt *DefaultToken) GetRowNum() int {

	return dt.lineNum
}

func (dt *DefaultToken) GetColumnNum() int {

	return dt.linePos
}

func (dt *DefaultToken) GetTokenKind() TokenKind {

	return dt.tk
}

func (dt *DefaultToken) GetName() *util.Name {
	panic("token_default 没有Name")
}

func (dt *DefaultToken) GetStringVal() string {

	return GetTokenString(dt.tk)
}

func (dt *DefaultToken) GetRadix() int {

	panic("implement me")
}

func (dt *DefaultToken) DebugToString() string {

	return fmt.Sprintf("关键字token: %v lineNum: %d pos: %d", dt.tk, dt.lineNum, dt.linePos)
}

func (dt *DefaultToken) CheckTokenKind() {

	if GetTokenKindTag(dt.tk) != TOKEN_TAG_DEFAULT {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_DEFAULT))
	}
}

func (dt *DefaultToken) Pos() int {

	return dt.pos
}

func (dt *DefaultToken) EndPos() int {

	return dt.endPos
}

func (dt *DefaultToken) GetSymbolTableIndex() int {

	return dt.inx
}

func (dt *DefaultToken) SetSymbolTableIndex(inx int) {

	dt.inx = inx
}
