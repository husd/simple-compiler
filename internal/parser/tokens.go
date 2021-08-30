package parser

import "husd.com/v0/util"

/**
 * 这里定义Java的词法分析器的所有的token
 */
type Tokens struct {
	tokenName []*util.Name // The names of all Tokens.
	key       []*tokenKind //Keyword array. Maps name indices to Token.
	maxKey    int          // key的最大索引
}

func NewTokens(c *util.Context) *Tokens {

	tks := Tokens{}

	//初始化所有的关键字

	return &tks
}

//这个是根据Name，返回是关键字 还是标识符 还是什么其它的
func lookupTokenKind(n *util.Name) *tokenKind {

	//TODO husd
	return TOKEN_KIND_IDENTIFIER
}
