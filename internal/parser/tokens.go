package parser

import "husd.com/v0/util"

/**
 * 这里定义Java的词法分析器的所有的token
 * @author hushengdong
 */
type Tokens struct {
	tokenName []*util.Name    // The names of all Tokens.
	key       [114]*tokenKind //Keyword array. Maps name indices to Token.
	maxKey    int             // key的最大索引

	keyword map[string]*tokenKind
}

func NewTokens(c *util.Context) *Tokens {

	tks := Tokens{}

	//初始化所有的关键字
	tks.keyword = make(map[string]*tokenKind)
	// 先这么用hashmap解决，后续再考虑优化成数组解决 TODO husd
	tks.keyword[TOKEN_KIND_EOF.Name] = TOKEN_KIND_EOF
	tks.keyword[TOKEN_KIND_ERROR.Name] = TOKEN_KIND_ERROR
	tks.keyword[TOKEN_KIND_IDENTIFIER.Name] = TOKEN_KIND_IDENTIFIER
	tks.keyword[TOKEN_KIND_ABSTRACT.Name] = TOKEN_KIND_ABSTRACT
	tks.keyword[TOKEN_KIND_ASSERT.Name] = TOKEN_KIND_ASSERT
	tks.keyword[TOKEN_KIND_BOOLEAN.Name] = TOKEN_KIND_BOOLEAN
	tks.keyword[TOKEN_KIND_BREAK.Name] = TOKEN_KIND_BREAK
	tks.keyword[TOKEN_KIND_BYTE.Name] = TOKEN_KIND_BYTE
	tks.keyword[TOKEN_KIND_CASE.Name] = TOKEN_KIND_CASE
	tks.keyword[TOKEN_KIND_CATCH.Name] = TOKEN_KIND_CATCH
	tks.keyword[TOKEN_KIND_CHAR.Name] = TOKEN_KIND_CHAR
	tks.keyword[TOKEN_KIND_CLASS.Name] = TOKEN_KIND_CLASS
	tks.keyword[TOKEN_KIND_CONST.Name] = TOKEN_KIND_CONST
	tks.keyword[TOKEN_KIND_CONTINUE.Name] = TOKEN_KIND_CONTINUE
	tks.keyword[TOKEN_KIND_DEF.Name] = TOKEN_KIND_DEF
	tks.keyword[TOKEN_KIND_DO.Name] = TOKEN_KIND_DO
	tks.keyword[TOKEN_KIND_DOUBLE.Name] = TOKEN_KIND_DOUBLE
	tks.keyword[TOKEN_KIND_ELSE.Name] = TOKEN_KIND_ELSE
	tks.keyword[TOKEN_KIND_ENUM.Name] = TOKEN_KIND_ENUM
	tks.keyword[TOKEN_KIND_EXTENDS.Name] = TOKEN_KIND_EXTENDS
	tks.keyword[TOKEN_KIND_FINAL.Name] = TOKEN_KIND_FINAL
	tks.keyword[TOKEN_KIND_FINALLY.Name] = TOKEN_KIND_FINALLY
	tks.keyword[TOKEN_KIND_FLOAT.Name] = TOKEN_KIND_FLOAT
	tks.keyword[TOKEN_KIND_FOR.Name] = TOKEN_KIND_FOR
	tks.keyword[TOKEN_KIND_GOTO.Name] = TOKEN_KIND_GOTO
	tks.keyword[TOKEN_KIND_IF.Name] = TOKEN_KIND_IF
	tks.keyword[TOKEN_KIND_IMPLEMENTS.Name] = TOKEN_KIND_IMPLEMENTS
	tks.keyword[TOKEN_KIND_IMPORT.Name] = TOKEN_KIND_IMPORT
	tks.keyword[TOKEN_KIND_INSTANCEOF.Name] = TOKEN_KIND_INSTANCEOF
	tks.keyword[TOKEN_KIND_INT.Name] = TOKEN_KIND_INT
	tks.keyword[TOKEN_KIND_INTERFACE.Name] = TOKEN_KIND_INTERFACE
	tks.keyword[TOKEN_KIND_LONG.Name] = TOKEN_KIND_LONG
	tks.keyword[TOKEN_KIND_NATIVE.Name] = TOKEN_KIND_NATIVE
	tks.keyword[TOKEN_KIND_NEW.Name] = TOKEN_KIND_NEW
	tks.keyword[TOKEN_KIND_PACKAGE.Name] = TOKEN_KIND_PACKAGE
	tks.keyword[TOKEN_KIND_PRIVATE.Name] = TOKEN_KIND_PRIVATE
	tks.keyword[TOKEN_KIND_PROTECTED.Name] = TOKEN_KIND_PROTECTED
	tks.keyword[TOKEN_KIND_PUBLIC.Name] = TOKEN_KIND_PUBLIC
	tks.keyword[TOKEN_KIND_RETURN.Name] = TOKEN_KIND_RETURN
	tks.keyword[TOKEN_KIND_SHORT.Name] = TOKEN_KIND_SHORT
	tks.keyword[TOKEN_KIND_STATIC.Name] = TOKEN_KIND_STATIC
	tks.keyword[TOKEN_KIND_STRICTFP.Name] = TOKEN_KIND_STRICTFP
	tks.keyword[TOKEN_KIND_SUPER.Name] = TOKEN_KIND_SUPER
	tks.keyword[TOKEN_KIND_SWITCH.Name] = TOKEN_KIND_SWITCH
	tks.keyword[TOKEN_KIND_SYNCHRONIZED.Name] = TOKEN_KIND_SYNCHRONIZED
	tks.keyword[TOKEN_KIND_THIS.Name] = TOKEN_KIND_THIS
	tks.keyword[TOKEN_KIND_THROW.Name] = TOKEN_KIND_THROW
	tks.keyword[TOKEN_KIND_THROWS.Name] = TOKEN_KIND_THROWS
	tks.keyword[TOKEN_KIND_TRANSIENT.Name] = TOKEN_KIND_TRANSIENT
	tks.keyword[TOKEN_KIND_TRY.Name] = TOKEN_KIND_TRY
	tks.keyword[TOKEN_KIND_VOID.Name] = TOKEN_KIND_VOID
	tks.keyword[TOKEN_KIND_VOLATILE.Name] = TOKEN_KIND_VOLATILE
	tks.keyword[TOKEN_KIND_WHILE.Name] = TOKEN_KIND_WHILE
	tks.keyword[TOKEN_KIND_INT_LITERAL.Name] = TOKEN_KIND_INT_LITERAL
	tks.keyword[TOKEN_KIND_LONG_LITERAL.Name] = TOKEN_KIND_LONG_LITERAL
	tks.keyword[TOKEN_KIND_FLOAT_LITERAL.Name] = TOKEN_KIND_FLOAT_LITERAL
	tks.keyword[TOKEN_KIND_DOUBLE_LITERAL.Name] = TOKEN_KIND_DOUBLE_LITERAL
	tks.keyword[TOKEN_KIND_CHAR_LITERAL.Name] = TOKEN_KIND_CHAR_LITERAL
	tks.keyword[TOKEN_KIND_STRING_LITERAL.Name] = TOKEN_KIND_STRING_LITERAL
	tks.keyword[TOKEN_KIND_TRUE.Name] = TOKEN_KIND_TRUE
	tks.keyword[TOKEN_KIND_FALSE.Name] = TOKEN_KIND_FALSE
	tks.keyword[TOKEN_KIND_NULL.Name] = TOKEN_KIND_NULL
	tks.keyword[TOKEN_KIND_UNDERSCORE.Name] = TOKEN_KIND_UNDERSCORE
	tks.keyword[TOKEN_KIND_ARROW.Name] = TOKEN_KIND_ARROW
	tks.keyword[TOKEN_KIND_COLCOL.Name] = TOKEN_KIND_COLCOL
	tks.keyword[TOKEN_KIND_LPAREN.Name] = TOKEN_KIND_LPAREN
	tks.keyword[TOKEN_KIND_RPAREN.Name] = TOKEN_KIND_RPAREN
	tks.keyword[TOKEN_KIND_LBRACE.Name] = TOKEN_KIND_LBRACE
	tks.keyword[TOKEN_KIND_RBRACE.Name] = TOKEN_KIND_RBRACE
	tks.keyword[TOKEN_KIND_LBRACKET.Name] = TOKEN_KIND_LBRACKET
	tks.keyword[TOKEN_KIND_RBRACKET.Name] = TOKEN_KIND_RBRACKET
	tks.keyword[TOKEN_KIND_SEMI.Name] = TOKEN_KIND_SEMI
	tks.keyword[TOKEN_KIND_COMMA.Name] = TOKEN_KIND_COMMA
	tks.keyword[TOKEN_KIND_DOT.Name] = TOKEN_KIND_DOT
	tks.keyword[TOKEN_KIND_ELLIPSIS.Name] = TOKEN_KIND_ELLIPSIS
	tks.keyword[TOKEN_KIND_EQ.Name] = TOKEN_KIND_EQ
	tks.keyword[TOKEN_KIND_GT.Name] = TOKEN_KIND_GT
	tks.keyword[TOKEN_KIND_LT.Name] = TOKEN_KIND_LT
	tks.keyword[TOKEN_KIND_BANG.Name] = TOKEN_KIND_BANG
	tks.keyword[TOKEN_KIND_TILDE.Name] = TOKEN_KIND_TILDE
	tks.keyword[TOKEN_KIND_QUES.Name] = TOKEN_KIND_QUES
	tks.keyword[TOKEN_KIND_COLON.Name] = TOKEN_KIND_COLON
	tks.keyword[TOKEN_KIND_EQEQ.Name] = TOKEN_KIND_EQEQ
	tks.keyword[TOKEN_KIND_LTEQ.Name] = TOKEN_KIND_LTEQ
	tks.keyword[TOKEN_KIND_GTEQ.Name] = TOKEN_KIND_GTEQ
	tks.keyword[TOKEN_KIND_BANGEQ.Name] = TOKEN_KIND_BANGEQ
	tks.keyword[TOKEN_KIND_AMPAMP.Name] = TOKEN_KIND_AMPAMP
	tks.keyword[TOKEN_KIND_BARBAR.Name] = TOKEN_KIND_BARBAR
	tks.keyword[TOKEN_KIND_PLUSPLUS.Name] = TOKEN_KIND_PLUSPLUS
	tks.keyword[TOKEN_KIND_SUBSUB.Name] = TOKEN_KIND_SUBSUB
	tks.keyword[TOKEN_KIND_PLUS.Name] = TOKEN_KIND_PLUS
	tks.keyword[TOKEN_KIND_SUB.Name] = TOKEN_KIND_SUB
	tks.keyword[TOKEN_KIND_STAR.Name] = TOKEN_KIND_STAR
	tks.keyword[TOKEN_KIND_SLASH.Name] = TOKEN_KIND_SLASH
	tks.keyword[TOKEN_KIND_AMP.Name] = TOKEN_KIND_AMP
	tks.keyword[TOKEN_KIND_BAR.Name] = TOKEN_KIND_BAR
	tks.keyword[TOKEN_KIND_CARET.Name] = TOKEN_KIND_CARET
	tks.keyword[TOKEN_KIND_PERCENT.Name] = TOKEN_KIND_PERCENT
	tks.keyword[TOKEN_KIND_LTLT.Name] = TOKEN_KIND_LTLT
	tks.keyword[TOKEN_KIND_GTGT.Name] = TOKEN_KIND_GTGT
	tks.keyword[TOKEN_KIND_GTGTGT.Name] = TOKEN_KIND_GTGTGT
	tks.keyword[TOKEN_KIND_PLUSEQ.Name] = TOKEN_KIND_PLUSEQ
	tks.keyword[TOKEN_KIND_SUBEQ.Name] = TOKEN_KIND_SUBEQ
	tks.keyword[TOKEN_KIND_STAREQ.Name] = TOKEN_KIND_STAREQ
	tks.keyword[TOKEN_KIND_SLASHEQ.Name] = TOKEN_KIND_SLASHEQ
	tks.keyword[TOKEN_KIND_AMPEQ.Name] = TOKEN_KIND_AMPEQ
	tks.keyword[TOKEN_KIND_BAREQ.Name] = TOKEN_KIND_BAREQ
	tks.keyword[TOKEN_KIND_CARETEQ.Name] = TOKEN_KIND_CARETEQ
	tks.keyword[TOKEN_KIND_PERCENTEQ.Name] = TOKEN_KIND_PERCENTEQ
	tks.keyword[TOKEN_KIND_LTLTEQ.Name] = TOKEN_KIND_LTLTEQ
	tks.keyword[TOKEN_KIND_GTGTEQ.Name] = TOKEN_KIND_GTGTEQ
	tks.keyword[TOKEN_KIND_GTGTGTEQ.Name] = TOKEN_KIND_GTGTGTEQ
	tks.keyword[TOKEN_KIND_MONKEYS_AT.Name] = TOKEN_KIND_MONKEYS_AT
	tks.keyword[TOKEN_KIND_CUSTOM.Name] = TOKEN_KIND_CUSTOM

	return &tks
}

//这个是根据Name，返回是关键字 还是标识符 还是什么其它的
func (ts *Tokens) lookupTokenKind(n *util.Name) *tokenKind {

	//是关键字，就返回，否则就是一个标识符
	if tk, ok := ts.keyword[n.NameStr]; ok {
		return tk
	}
	return TOKEN_KIND_IDENTIFIER
}
