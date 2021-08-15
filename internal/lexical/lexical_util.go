package lexical

/**
 * 这个类，提供了一些公共的方法 和具体的词法分析器无关的代码
 * 词法分析器可能会做很多事情
 */

//空或者换行，都是停止符号
func endChar(ch uint8) bool {

	return ch == ' ' || ch == '\n'
}
