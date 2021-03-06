package ast_tree2

import "husd.com/v0/util"

/**
 * For example:
 * <pre>
 *   <em>expression</em> . <em>identifier</em>
 * </pre>
 * @author hushengdong
 */
type MemberSelectTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	MemberSelectTreeV2_()

	// --
	GetExpression() ExpressionTreeV2
	GetIdentifier() *util.Name
}
