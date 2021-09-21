package ast_tree2

/**
 * For example:
 * <pre>
 *   case <em>expression</em> :
 *       <em>statements</em>
 *
 *   default :
 *       <em>statements</em>
 * </pre>
 * @author hushengdong
 */
type CaseTreeV2 interface {
	TreeType() *TreeType
	//--
	GetExpression() ExpressionTreeV2
	GetStatements() StatementTreeV2
}
