package ast_tree2

/**
 * For example:
 * <pre>{@code
 *   ()->{}
 *   (List<String> ls)->ls.size()
 *   (x,y)-> { return x + y; }
 * }</pre>
 *
 * @author hushengdong
 */
type LambdaExpressionTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	LambdaExpressionTreeV2_()

	// ---
	GetParameters() *[]VariableTreeV2
	GetBody() TreeV2
	GetBodyKind() LambdaExpressionTreeV2BodyKind
}

type LambdaExpressionTreeV2BodyKind string

const (
	EXPRESSION = "expression"
	STATEMENT  = "statement"
)
