package parser

/**
 * 主要定义了解析词法分析器的解析文本功能的主要方法
 * 先抽象一下，解析一个文本文件，都需要哪些逻辑。
 * @author hushengdong
 */
type lexer interface {
	/**
	 * 移动指针到下一个 token
	 */
	NextToken()
	/**
	 * 返回当前指针指向的 token
	 */
	Token() token

	Ahead(len int) token

	PreToken() token

	ErrPos() int

	SetErrPos(pos int)

	GetLineMap() *lineMap
}
