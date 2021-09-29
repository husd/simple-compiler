package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>type</em> &lt; <em>typeArguments</em> &gt;
 * </pre>
 * List<Integer> 这样类似的语法
 * @author hushengdong
 */
type ParameterizedTypeTreeV2 interface {
	GetTreeType() TreeType
	ParameterizedTypeTreeV2_()
	// --
	getType() TreeV2
	getTypeArguments() *[]TreeV2
}
