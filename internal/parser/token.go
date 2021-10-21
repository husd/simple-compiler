package parser

import (
	"husd.com/v0/util"
)

// Token 词法分析器解析出来的最小单元
type Token interface {
	/**
	 * token的类型
	 */
	GetTokenKind() TokenKind
	/**
	 * token的属性
	 */
	GetName() *util.Name
	/**
	 * 当前token的字面量是什么
	 */
	GetStringVal() string
	/**
	 * 数字类型的token会有
	 */
	GetRadix() int
	/**
	 * 检查当前token的TAG是否正确
	 */
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
	/**
	 * 在源代码中的行号
	 */
	GetRowNum() int
	/**
	 * 在源代码中的列
	 */
	GetColumnNum() int
}
