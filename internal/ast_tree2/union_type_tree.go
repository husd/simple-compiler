package ast_tree2

/**
 *
 * @author hushengdong
 */
type UnionTypeTreeV2 interface {
	TreeType() *TreeType
	UnionTypeTreeV2_()
	// --
	GetTypeAlternatives() *[]TreeV2
}
