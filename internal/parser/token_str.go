package parser

import (
	"fmt"
	"husd.com/v0/util"
)

type stringToken struct {
	tk    *tokenKind
	start int
	end   int

	val string //字面量
}

func newStringToken(tk *tokenKind, start int, end int, val string) *stringToken {

	res := stringToken{tk, start, end, val}
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

func (st *stringToken) CheckTokenKind() {
	if st.tk.Tag != TOKEN_TAG_STRING {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_STRING))
	}
}
