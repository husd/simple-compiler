package parser

/**
 * 这个类是判断Java标识符的
 * @author hushengdong
 */

/**
 * 是否是可以忽略的字符 有些字符在Java里是直接忽略的，目前有2类：
 * 一、ISO的一些控制符号，是直接忽略的，这些字符不是空格
 * '\u0000' through '\u0008'
 * '\u000E' through '\u001B'
 * '\u007F' through '\u009F'
 *
 * 二、all characters that have the FORMAT general category value
 */
func isIdentifierIgnorable(ch rune) bool {

	if ch >= '\u0000' && ch <= '\u0008' {
		return true
	}
	if ch >= '\u000E' && ch <= '\u001B' {
		return true
	}
	if ch >= '\u007F' && ch <= '\u009F' {
		return true
	}
	//其它的可以忽略的，暂时先不管 TODO husd
	return false
}

func isJavaIdentifierPart(ch rune) bool {

	//这块涉及到复杂的unicode编码问题 所以暂时先返回true，表示认可，都可以是Java的标识符
	return true
}
