package ast_tree2

/**
 *
 * @author hushengdong
 */
type LabeledStatementTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()
	LabeledStatementTreeV2_()

	// --
	GetLabel() string
	GetStatement() StatementTreeV2
}
