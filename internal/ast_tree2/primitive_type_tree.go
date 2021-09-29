package ast_tree2

import "husd.com/v0/lang"

/**
 * For example:
 * <pre>
 *   <em>primitiveTypeKind</em>
 * </pre>
 * @author hushengdong
 */
type PrimitiveTypeTreeV2 interface {
	GetTreeType() TreeType
	PrimitiveTypeTreeV2_()
	// --
	GetPrimitiveTypeKind() lang.TypeKind
}
