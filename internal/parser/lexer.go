package parser

/**
 * 主要定义了解析词法分析器的解析文本功能的主要方法
 * 先抽象一下，解析一个文本文件，都需要哪些逻辑。
 * @author hushengdong
 */
type lexer interface {
	/**
	 * 移动指针到下一个 Token
	 */
	NextToken()
	/**
	 * 返回当前指针指向的 Token
	 */
	Token() Token
	/**
	 * 向前看几个token 预读后面的token
	 */
	LookAheadByIndex(lookahead int) Token
	/**
	 * 前面1个token
	 */
	PreToken() Token

	ErrPos() int

	SetErrPos(pos int)

	GetLineMap() *lineMap
}
