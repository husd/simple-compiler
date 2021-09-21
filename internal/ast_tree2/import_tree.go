package ast_tree2

/**
 * For example:
 * <pre>
 *   import <em>qualifiedIdentifier</em> ;
 *
 *   static import <em>qualifiedIdentifier</em> ;
 * </pre>
 * @author hushengdong
 */
type ImportTreeV2 interface {
	TreeType() *TreeType
	ImportTreeV2_()
	// --
	IsStatic() bool
	GetQualifiedIdentifier() TreeV2
}
