package ast_tree2

/**
 * For example:
 * <pre>
 *   try
 *       <em>block</em>
 *   <em>catches</em>
 *   finally
 *       <em>finallyBlock</em>
 * </pre>
 * @author hushengdong
 */
type TryTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()
	TryTreeV2_()
	// --
	GetBlock() BlockTreeV2
	GetCatches() *[]CatchTreeV2
	GetFinallyBlock() BlockTreeV2
	GetResources() *[]TreeV2
}
