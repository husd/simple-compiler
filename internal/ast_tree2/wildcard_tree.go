package ast_tree2

/**
 * For example:
 * <pre>
 *   ?
 *
 *   ? extends <em>bound</em>
 *
 *   ? super <em>bound</em>
 * </pre>
 * @author hushengdong
 */
type WildcardTreeV2 interface {
	TreeType() *TreeType
	WildcardTreeV2_()
	// --
	GetBound() TreeV2
}
