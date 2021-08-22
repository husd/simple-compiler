package parser

// Token 词法分析器解析出来的最小单元
type Token struct {
	TokenKind tokenKind
	StartPos  int
	EndPos    int
}

type namedToken struct {
	Token Token
	name  name
}
