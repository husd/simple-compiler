package ast_tree2

/**
 *
 * @author hushengdong
 */
type UnionTypeTreeV2 interface {
	GetTreeType() TreeType
	UnionTypeTreeV2_()
	// --
	GetTypeAlternatives() *[]TreeV2
}
