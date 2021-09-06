package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>name</em>
 *
 *   <em>name</em> extends <em>bounds</em>
 *
 *   <em>annotations</em> <em>name</em>
 * </pre>
 * @author hushengdong
 */
type TypeParameterTreeV2 interface {
	TreeType() TreeType
	TypeParameterTreeV2_()

	// --
	GetName() string
	GetBounds() TreeV2
	GetAnnotations() AnnotationTreeV2 //
}
