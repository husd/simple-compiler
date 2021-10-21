package parser

import (
	"fmt"
	"husd.com/v0/code"
)

/**
 * 树的工厂类
 * @author hushengdong
 */

var empty_tree_node *TreeNode
var empty_tree_node_arr *([]TreeNode)

func init() {

	fmt.Println("ast tree maker......")
	initEmpty()
	temp := make([]TreeNode, 0, 0)
	empty_tree_node_arr = &temp
}

// 临时写个空的树节点
func GetEmptyTreeNode() *TreeNode {

	return empty_tree_node
}

func GetEmptyTreeNodeArray() *([]TreeNode) {

	return empty_tree_node_arr
}

func initEmpty() {

	res := &TreeNode{}
	res.index = -1
	res.tag = Tree_node_tag_skip
	res.childrenCount = 0
	res.name = "empty"
	res.expr_or_state = node_type_unknown

	empty_tree_node = res
}

// error node
func NewErrorTreeNode(pos int, msg string) *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = Tree_node_tag_erroneous
	res.childrenCount = 0
	res.name = "error:" + msg
	res.expr_or_state = node_type_unknown
	res.pos = pos

	return res
}

func NewDummyTreeNode() *TreeNode {

	res := &TreeNode{}
	res.index = -1
	res.tag = Tree_node_tag_erroneous
	res.childrenCount = 0
	res.children = make([]*TreeNode, 2, 2)
	res.name = "dummy node"
	res.expr_or_state = node_type_unknown

	return res
}

func NewBlockTreeNode(token Token) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = Tree_node_tag_block
	res.childrenCount = 0
	res.children = make([]*TreeNode, 5, 5)
	res.name = "block node"
	res.expr_or_state = node_type_unknown

	return res
}

/**
 * if 有3个节点 condition truePart falsePart
 */
func NewIfTreeNode(token Token) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = Tree_node_tag_if
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "if node"
	res.expr_or_state = node_type_statement

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

	return res
}

func NewLiteralTreeNode(token Token, typeTag *code.TypeTag, val interface{}) *TreeNode {

	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = Tree_node_tag_literal
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "字面量"
	res.expr_or_state = node_type_statement

	res.typeTag = typeTag
	res.val = val

	return res
}

func NewIdentifyTreeNode(token Token) *TreeNode {

	// TODO
	res := &TreeNode{}
	res.index = token.GetSymbolTableIndex()
	res.tag = Tree_node_tag_ident
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "标识符"
	res.expr_or_state = node_type_statement

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

	return res
}
