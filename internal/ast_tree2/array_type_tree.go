package ast_tree2

/**
 * int[] arr 这样的定义了数组类型
 * @author hushengdong
 */
type ArrayTypeTreeV2 interface {
	TreeType() TreeType

	GetType() TreeV2
}
