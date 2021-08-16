package lexical

import "container/list"

type Lexer interface {
	//LexicalAnalysis
	/**
	 * 词法分析器接口，输入的是字符串形式的源文件代码，输出的是一段序列，表示
	 * 词素，Lexeme 结构
	 */
	LexicalAnalysis(str string) *list.List
}

// Token
/**
 * 这个是词法分析器的输出的最小单元，产生了类似这样的数据：
 * <token-name , attribute-value>
 */
type Token struct {
	//名字
	Name string
	// 在符号表里的位置
	Idx int
	//词素的类型
	Tag int
	// 在源代码的多少行
	LineNum int
}
