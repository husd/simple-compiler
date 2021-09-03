package util

/**
 * 这个类，提供了一些公共的方法 和具体的词法分析器无关的代码
 * 词法分析器可能会做很多事情
 * 这里的代码未来要删除，属于试水阶段的实验性质代码。 TODO delete later
 * @author hushengdong
 */

//空或者换行，都是停止符号
func EndChar(ch uint8) bool {

	return ch == ' ' || ch == '\n'
}

//空格
func BlankChar(ch uint8) bool {

	return ch == ' '
}

//换行
func EofChar(ch uint8) bool {

	return ch == '\n'
}
