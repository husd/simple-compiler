package parser

import "husd.com/v0/util"

/**
 * 这里定义Java的词法分析器的所有的token
 * @author hushengdong
 */
type Tokens struct {
	tokenName []*util.Name   // The names of all Tokens.
	key       [114]tokenKind // Keyword array. Maps name indices to Token.
	maxKey    int            // key的最大索引

	keyword map[string]tokenKind
}

func InstanceTokens(c *util.Context) *Tokens {

	ok, obj := c.Get(util.C_TOKENS)
	if ok {
		return obj.(*Tokens)
	}
	return NewTokens(c)
}

func NewTokens(c *util.Context) *Tokens {

	tks := &Tokens{}

	// 初始化所有的关键字
	keyword := make(map[string]tokenKind)
	for i := 2; i < len(token_kind_array); i++ {
		tk := token_kind_array[i]
		if tk == "" {
			continue
		}
		keyword[tk] = tokenKind(i)
	}

	tks.keyword = keyword
	c.Put(util.C_TOKENS, tks)
	return tks
}

//这个是根据Name，返回是关键字 还是标识符 还是什么其它的
func (ts *Tokens) lookupTokenKind(n *util.Name) tokenKind {

	// 是关键字，就返回，否则就是一个标识符
	if tk, ok := ts.keyword[n.NameStr]; ok {
		return tk
	}
	return TOKEN_KIND_IDENTIFIER
}
