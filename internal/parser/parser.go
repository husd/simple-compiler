package parser

import (
	"husd.com/v0/ast"
)

//Parser
/**
 * 语法分析接口，用来生成AST
 * 重要 这里有3个概念，非常重要 这几个概念都是语法树部分的内容
 * JCCompilationUnit
 * AbstractJCExpression 表达式 类似： a = a + 1;
 * AbstractJCStatement 声明 类似： String str = "123";
 */
type Parser interface {
	// Parse a compilation unit.
	ParseJCCompilationUnit() *ast.TreeNode
	//Parse an expression.
	ParseExpression() *ast.TreeNode
	//Parse a statement.
	ParseStatement() *ast.TreeNode
	//Parse a type.
	ParseType() *ast.TreeNode
}
