package ast_tree2

/**
 * {}
 * { statements }
 * static {statements}
 * extends StatementTreeV2
 * @author hushengdong
 */
type BlockTreeV2 interface {
	TreeType() *TreeType
	StatementTreeV2_()
	// --
	/**
	 * 其它语言可能没有static
	 */
	IsStatic() bool
	GetStatements() *[]StatementTreeV2
}
