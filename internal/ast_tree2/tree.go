package ast_tree2

/**
 * AST的抽象定义，这里的代码应该是和语言无关的，理论上可以轻松迁移到其它语言。
 * 所有的Tree都必须实现这个接口。
 * @author hushengdong
 */
type TreeV2 interface {
	/**
	 * 树的节点，有不同的类型
	 */
	TreeType() TreeType
}

/**
 * 这里定义了树的类型，这里为了更方便阅读代码，采用的是string
 */
type TreeType string

const (
	NODE_TYPE_PACKAGE = "声明package"
)

//树的节点
type AstTreeNode struct {
	nodeType TreeType
	msg      string //描述信息

}
