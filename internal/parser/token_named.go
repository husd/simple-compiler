package parser

import (
	"fmt"
	"husd.com/v0/util"
)

type namedToken struct {
	tk    *tokenKind
	start int
	end   int
	n     *util.Name
}

func newNamedToken(tk *tokenKind, start int, end int, n *util.Name) *namedToken {

	res := namedToken{tk, start, end, n}
	return &res
}

func (nt *namedToken) GetTokenKind() *tokenKind {

	return nt.tk
}

func (nt *namedToken) GetName() *util.Name {

	return nt.n
}

func (nt *namedToken) GetStringVal() string {

	return nt.n.NameStr
}

func (nt *namedToken) GetRadix() int {

	panic("implement me")
}

func (nt *namedToken) CheckTokenKind() {

	if nt.tk.Tag != TOKEN_TAG_NAMED {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_NAMED))
	}

}
