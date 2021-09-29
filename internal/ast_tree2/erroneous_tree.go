package ast_tree2

/**
 * 表示格式错误表达式的树节点。
 *
 * @author hushengdong
 */
type ErroneousTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()

	//--
	GetErrorTrees() *[]TreeV2
}
