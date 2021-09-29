package ast_tree2

/**
 * For example:
 * <pre>
 *    ;
 * </pre>
 * @author hushengdong
 */
type EmptyStatementTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()

	// --
	EmptyStatementTreeV2_()
}
