package ast_tree2

/**
 * statement tree是一个通用的表达式 不特指具体的类型
 *
 * @author hushengdong
 */
type StatementTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()
}
