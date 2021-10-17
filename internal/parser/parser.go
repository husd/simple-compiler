package parser

//Parser
/**
 * 语法分析接口，用来生成AST
 * 重要 这里有3个概念，非常重要 这几个概念都是语法树部分的内容
 * JCCompilationUnit
 * AbstractJCExpression 表达式 类似： a = a + 1;
 * AbstractJCStatement 声明 类似： String str = "123";
 *
 * statement ::== ; | expression; | if (expression) statement
 * ...带分号的 expression 就是 statement. 简言之，
 * expression 告诉解释器（编译器），这是可求值的，而 statement 则说请求值。
 *
 */
type Parser interface {
	// Parse a compilation unit.
	ParseJCCompilationUnit() *TreeNode
	//Parse an expression.
	ParseExpression() *TreeNode
	//Parse a statement.
	ParseStatement() *TreeNode
	//Parse a type.
	ParseType() *TreeNode
}
