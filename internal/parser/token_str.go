package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 字符串token
 */
type stringToken struct {
	tk *tokenKind

	lineNum int // 多少行
	linePos int // 位置

	val string //字面量

}

func newStringToken(tk *tokenKind, lineNum int, linePos int, val string) *stringToken {

	res := stringToken{tk, lineNum, linePos, val}
	return &res
}

func (st *stringToken) GetTokenKind() *tokenKind {

	return st.tk
}

func (st *stringToken) GetName() *util.Name {

	panic("implement me")
}

func (st *stringToken) GetStringVal() string {

	return st.val
}

func (st *stringToken) GetRadix() int {

	panic("implement me")
}

func (dt *stringToken) DebugToString() string {

	return fmt.Sprintf("stringtoken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (st *stringToken) CheckTokenKind() {
	if st.tk.Tag != TOKEN_TAG_STRING {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_STRING))
	}
}
