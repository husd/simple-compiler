package ast_tree2

/**
 * For example:
 * <pre>
 *   switch ( <em>expression</em> ) {
 *     <em>cases</em>
 *   }
 * </pre>
 * @author hushengdong
 */
type SwitchTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()
	SwitchTreeV2_()
	// --
	GetExpression() ExpressionTreeV2
	GetCases() *[]CaseTreeV2
}
