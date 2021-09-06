package ast_tree2

/**
 * For example:
 * <pre>
 *   synchronized ( <em>expression</em> )
 *       <em>block</em>
 * </pre>
 * @author hushengdong
 */
type SynchronizedTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()
	SynchronizedTreeV2_()
	// --
	GetExpression() ExpressionTreeV2
	GetBlock() BlockTreeV2
}
