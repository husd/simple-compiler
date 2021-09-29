package parser

import "husd.com/v0/util"

// Token 词法分析器解析出来的最小单元
type Token interface {
	/**
	 * token的类型
	 */
	GetTokenKind() tokenKind

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
	/**
	 * 设置token在符号表的位置
	 */
	SetSymbolTableIndex(index int)
}
