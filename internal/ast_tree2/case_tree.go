package ast_tree2

/**
 * switch a {
 *   case 10:
 *   default:
 * }
 * @author hushengdong
 */
type CaseTreeV2 interface {
	TreeType() TreeType
	//--
	GetExpression() ExpressionTreeV2
	GetStatements() StatementTreeV2
}
