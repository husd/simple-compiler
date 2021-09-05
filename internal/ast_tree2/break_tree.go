package ast_tree2

/**
 * break label
 * extends StatementTreeV2
 * @author hushengdong
 */
type BreakTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()

	// --
	/**
	 * 先返回string
	 */
	GetLabel() string
	BreakTreeV2_()
}
