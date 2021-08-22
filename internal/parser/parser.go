package parser

import tree "husd.com/v0/tree"

//Parser
/**
 * 重要 这里有3个概念，非常重要 这几个概念都是语法树部分的内容
 * JCCompilationUnit
 * JCExpression 表达式 类似： a = a + 1;
 * JCStatement 声明 类似： String str = "123";
 */
type Parser interface {
	// Parse a compilation unit.
	ParseJCCompilationUnit() tree.JCCompilationUnit
	//Parse an expression.
	ParseExpression() tree.JCExpression
	//Parse a statement.
	ParseStatement() tree.JCStatement
	//Parse a type.
	ParseType() tree.JCExpression
}
