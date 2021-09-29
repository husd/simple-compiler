package ast_tree2

/**
 * For example:
 * <pre>
 *   if ( <em>condition</em> )
 *      <em>thenStatement</em>
 *
 *   if ( <em>condition</em> )
 *       <em>thenStatement</em>
 *   else
 *       <em>elseStatement</em>
 * </pre>
 * @author hushengdong
 */
type IfTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()

	// --
	GetCondition() ExpressionTreeV2
	GetThenStatement() StatementTreeV2
	GetElseStatement() StatementTreeV2
}
