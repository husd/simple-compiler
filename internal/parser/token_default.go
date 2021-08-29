package parser

import (
	"fmt"
	"husd.com/v0/util"
)

type defaultToken struct {
	tk    *tokenKind
	start int
	end   int
}

func newDefaultToken(tk *tokenKind, start int, end int) *defaultToken {

	res := defaultToken{tk, start, end}
	return &res
}

func (dt *defaultToken) GetTokenKind() *tokenKind {
	return dt.tk
}

func (dt *defaultToken) GetName() *util.Name {
	panic("implement me")
}

func (dt *defaultToken) GetStringVal() string {
	panic("implement me")
}

func (dt *defaultToken) GetRadix() int {
	panic("implement me")
}

func (dt *defaultToken) CheckTokenKind() {

	if dt.tk.Tag != TOKEN_TAG_DEFAULT {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_DEFAULT))
	}
}
