package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/util"
)

/**
 * 树的工厂类
 * @author hushengdong
 */

var empty_tree_node *TreeNode
var empty_tree_node_arr *([]*TreeNode)

func init() {

	fmt.Println("ast tree maker......")
	initEmpty()
	temp := make([]*TreeNode, 0, 0)
	empty_tree_node_arr = &temp
}

// 临时写个空的树节点
func GetEmptyTreeNode() *TreeNode {

	return empty_tree_node
}

func GetEmptyTreeNodeArray() *([]*TreeNode) {

	return empty_tree_node_arr
}

func initEmpty() {

	res := &TreeNode{}
	res.index = -1
	res.tag = skip
	res.childrenCount = 0
	res.name = "empty"
	res.expr_or_state = node_type_unknown
	res.treeType = tt_nil

	empty_tree_node = res
}

// error node
func NewErrorTreeNode(pos int, msg string) *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = erroneous
	res.childrenCount = 0
	res.name = "error:" + msg
	res.expr_or_state = node_type_unknown
	res.pos = pos
	res.treeType = tt_erroneous

	return res
}

func NewDummyTreeNode() *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = erroneous
	res.childrenCount = 0
	res.children = make([]*TreeNode, 2, 2)
	res.name = "dummy node"
	res.expr_or_state = node_type_unknown
	res.treeType = tt_erroneous

	return res
}

func NewBlockTreeNode(token Token) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = block
	res.childrenCount = 0
	res.children = make([]*TreeNode, 5, 5)
	res.name = "block node"
	res.expr_or_state = node_type_unknown
	res.treeType = tt_block

	return res
}

/**
 * if 有3个节点 condition truePart falsePart
 */
func NewIfTreeNode(token Token) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = if_
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "if node"
	res.expr_or_state = node_type_statement
	res.treeType = tt_if

	return res
}

func NewCompareConditionTreeNode(token Token, tag TreeNodeTag) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = tag // 66 - 71 之间 ，这里暂且不考虑 TODO 函数调用
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "if node - condition"
	res.expr_or_state = node_type_statement
	res.treeType = tt_conditional_expression
	return res
}

func NewLiteralTreeNode(token Token, typeTag *code.TypeTag, val interface{}, tt TreeType) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = literal
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "字面量"
	res.expr_or_state = node_type_statement

	res.typeTag = typeTag
	res.val = val

	res.treeType = tt //这个是动态计算的
	return res
}

func NewIdentifyTreeNode(token Token, name *util.Name, sym *Symbol) *TreeNode {

	// TODO
	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = ident
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "标识符"
	res.expr_or_state = node_type_statement
	res.treeType = tt_identifier
	res.n = name
	res.symbol = sym

	return res
}

func NewUnaryTreeNode(tag TreeNodeTag) *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = tag
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "一元操作符"
	res.expr_or_state = node_type_statement
	res.treeType = getTreeTypeByTreeNodeTag(tag)
	return res
}

func NewBinaryOpTreeNode(tag TreeNodeTag) *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = tag
	res.childrenCount = 2
	res.children = make([]*TreeNode, 2, 2)
	res.name = "二元操作符"
	res.expr_or_state = node_type_expression
	res.treeType = getTreeTypeByTreeNodeTag(tag)
	return res
}

/**
 * 注意要和字面量区分开来
 */
func NewPrimitiveTypeTree(index int, tt *code.TypeTag) *TreeNode {

	res := &TreeNode{}
	res.index = index
	res.tag = typeident // 表示定义表示符号的类型 是基本类型 Java里有2种: 基本类型 和 Class
	res.childrenCount = 0
	//res.children = make([]*TreeNode, 2, 2)
	res.name = "基本类型定义"
	res.expr_or_state = node_type_expression
	// 不需要val
	res.typeTag = tt

	res.treeType = tt_primitive_type

	return res
}

/**
 * 数组类型的 例如： int[] arr = {1,2,3};
 * 这里用来表示 int[] 这个元素，表示一个数组类型
 */
func NewTypeArrayTreeNode(t *TreeNode) *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = typearray
	res.childrenCount = 0
	//res.children = make([]*TreeNode, 2, 2)
	res.name = "数组定义"
	res.expr_or_state = node_type_expression

	res.val = t // t表示数组的类型 ，例如 int[] 是 int 类型的数组
	res.treeType = tt_array_type

	return res
}
