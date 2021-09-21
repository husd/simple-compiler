package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>expression</em> # <em>[ identifier | new ]</em>
 * </pre>
 * @author hushengdong
 */
type MemberReferenceTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()
	MemberReferenceTreeV2_()

	// ---
	GetMode() ReferenceMode
	GetQualifierExpression() ExpressionTreeV2
	GetName() string
	GetTypeArguments() *[]ExpressionTreeV2
}

type ReferenceMode string

const (
	INVOKE = "invoke"
	NEW    = "new"
)
