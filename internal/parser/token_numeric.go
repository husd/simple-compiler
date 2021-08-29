package parser

import (
	"fmt"
	"husd.com/v0/util"
)

type numericToken struct {
	tk    *tokenKind
	start int
	end   int

	val   string
	radix int
}

func newNumericToken(tk *tokenKind, start int, end int, val string, radix int) *numericToken {

	res := numericToken{tk, start, end, val, radix}
	return &res
}

func (nt *numericToken) GetTokenKind() *tokenKind {
	return nt.tk
}

func (nt *numericToken) GetName() *util.Name {
	panic("implement me")
}

func (nt *numericToken) GetStringVal() string {
	panic("implement me")
}

func (nt *numericToken) GetRadix() int {
	return nt.radix
}

func (nt *numericToken) CheckTokenKind() {
	if nt.tk.Tag != TOKEN_TAG_NUMERIC {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_NUMERIC))
	}
}
