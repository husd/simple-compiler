package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 这个是java的关键字和运算符号之类
 */
type defaultToken struct {
	tk *tokenKind

	lineNum int // 多少行
	linePos int // 位置

}

func newDefaultToken(tk *tokenKind, lineNum int, linePos int) *defaultToken {

	res := defaultToken{tk, lineNum, linePos}
	return &res
}

func (dt *defaultToken) GetTokenKind() *tokenKind {

	return dt.tk
}

func (dt *defaultToken) GetName() *util.Name {

	n := util.Name{NameStr: dt.tk.Name, Index: 0}
	return &n
}

func (dt *defaultToken) GetStringVal() string {

	return dt.tk.Name
}

func (dt *defaultToken) GetRadix() int {

	panic("implement me")
}

func (dt *defaultToken) DebugToString() string {

	return fmt.Sprintf("关键字token: %v lineNum: %d pos: %d", dt.tk.Name, dt.lineNum, dt.linePos)
}

func (dt *defaultToken) CheckTokenKind() {

	if dt.tk.Tag != TOKEN_TAG_DEFAULT {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_DEFAULT))
	}
}
