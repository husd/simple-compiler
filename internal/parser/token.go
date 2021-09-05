package parser

import "husd.com/v0/util"

type tokenTag int

const (
	TOKEN_TAG_DEFAULT tokenTag = 1
	TOKEN_TAG_NAMED   tokenTag = 2
	TOKEN_TAG_STRING  tokenTag = 3
	TOKEN_TAG_NUMERIC tokenTag = 4
)

// Token 词法分析器解析出来的最小单元
type Token interface {
	GetTokenKind() *tokenKind

	GetName() *util.Name

	GetStringVal() string
	/**
	 * 数字类型的token会有
	 */
	GetRadix() int

	CheckTokenKind()
	/**
	 * DEBUG使用的
	 */
	DebugToString() string
	/**
	 * position started in the source file
	 */
	Pos() int
	/**
	 * position ended in the source file
	 */
	EndPos() int
	/**
	 * 在符号表里的位置
	 */
	GetSymbolTableIndex() int

	SetSymbolTableIndex(index int)
}
