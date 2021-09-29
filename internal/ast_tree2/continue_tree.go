package ast_tree2

/**
 * For example:
 * <pre>
 *   continue;
 *   continue <em>label</em> ;
 * </pre>
 * @author hushengdong
 */
type ContinueTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()

	// --
	GetLabel() string
	ContinueTreeV2_()
}
