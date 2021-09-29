package ast_tree2

import "husd.com/v0/lang"

/**
 * 修饰符之类 private static final and so on
 * @author hushengdong
 */
type ModifiersTreeV2 interface {
	GetTreeType() TreeType

	// --
	GetFlags() *[]lang.Modifier
	ModifiersTreeV2_()
}
