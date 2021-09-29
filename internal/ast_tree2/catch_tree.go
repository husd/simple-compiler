package ast_tree2

/**
 * try catch  extends Tree
 * For example:
 * <pre>
 *   catch ( <em>parameter</em> )
 *       <em>block</em>
 * </pre>
 * @author hushengdong
 */
type CatchTreeV2 interface {
	GetTreeType() TreeType

	GetParameter() VariableTreeV2
	GetBlock() BlockTreeV2
}
