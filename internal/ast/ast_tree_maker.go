package ast

import "fmt"

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
	res.tag = Tree_node_tag_skip
	res.childrenCount = 0
	res.name = "empty"
	res.expr_or_state = node_type_unknown

	empty_tree_node = res
}

// error node
func NewErrorTreeNode(msg string) *TreeNode {

	res := &TreeNode{}
	res.tag = Tree_node_tag_erroneous
	res.childrenCount = 0
	res.name = "error:" + msg
	res.expr_or_state = node_type_unknown

	return res
}

func NewDummyTreeNode() *TreeNode {

	res := &TreeNode{}
	res.tag = Tree_node_tag_erroneous
	res.childrenCount = 0
	res.children = make([]*TreeNode, 2, 2)
	res.name = "dummy node"
	res.expr_or_state = node_type_unknown

	return res
}

func NewBlockTreeNode() *TreeNode {

	res := &TreeNode{}
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
func NewIfTreeNode() *TreeNode {

	res := &TreeNode{}
	res.tag = Tree_node_tag_if
	res.childrenCount = 0
	res.children = make([]*TreeNode, 3, 3)
	res.name = "if node"
	res.expr_or_state = node_type_statement

	return res
}
