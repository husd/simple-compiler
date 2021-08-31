package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * true false enum void this super byte short char int long float double等
 * 这些都统归为 namedToken 表示的是定义变量的属性的，可以看到大部分都是基本类型
 */
type namedToken struct {
	tk      *tokenKind
	lineNum int // 多少行
	linePos int // 位置
	n       *util.Name
}

func newNamedToken(tk *tokenKind, lineNum int, linePos int, n *util.Name) *namedToken {

	res := namedToken{tk, lineNum, linePos, n}
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

func (dt *namedToken) DebugToString() string {

	return fmt.Sprintf("namedToken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (nt *namedToken) CheckTokenKind() {

	if nt.tk.Tag != TOKEN_TAG_NAMED {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_NAMED))
	}

}
