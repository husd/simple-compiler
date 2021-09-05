package ast_tree2

/**
 * For example:
 * <pre>
 *    ;
 * </pre>
 * @author hushengdong
 */
type EmptyStatementTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()

	// --
	EmptyStatementTreeV2_()
}
