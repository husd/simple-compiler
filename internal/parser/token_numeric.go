package parser

import (
	"fmt"
	"husd.com/v0/util"
)

/**
 * 数字类的token
 * @author hushengdong
 */
type numericToken struct {
	tk      *tokenKind
	lineNum int // 多少行
	linePos int // 位置

	val   string
	radix int
}

func newNumericToken(tk *tokenKind, lineNum int, linePos int, val string, radix int) *numericToken {

	res := numericToken{tk, lineNum, linePos, val, radix}
	return &res
}

func (nt *numericToken) GetTokenKind() *tokenKind {
	return nt.tk
}

func (nt *numericToken) GetName() *util.Name {

	n := util.Name{NameStr: nt.val, Index: 0}
	return &n
}

func (nt *numericToken) GetStringVal() string {

	return nt.val
}

func (nt *numericToken) GetRadix() int {
	return nt.radix
}

func (dt *numericToken) DebugToString() string {

	return fmt.Sprintf("numericToken: %v lineNum: %d pos: %d", dt.GetStringVal(), dt.lineNum, dt.linePos)
}

func (nt *numericToken) CheckTokenKind() {
	if nt.tk.Tag != TOKEN_TAG_NUMERIC {
		panic(fmt.Sprintf("错误的token kind ，应该是：%d", TOKEN_TAG_NUMERIC))
	}
}
