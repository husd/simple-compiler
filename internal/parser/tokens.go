package parser

import "husd.com/v0/util"

/**
 * 这里定义Java的词法分析器的所有的token
 */
type tokens struct {
	name string
}

//这个是根据Name，返回是关键字 还是标识符 还是什么其它的
func lookupTokenKind(*util.Name) *tokenKind {

	//TODO husd
	return TOKEN_KIND_ERROR
}
