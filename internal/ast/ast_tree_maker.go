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
func EmptyTreeNode() *TreeNode {

	return empty_tree_node
}

func EmptyTreeNodeArray() *([]TreeNode) {

	return empty_tree_node_arr
}

func initEmpty() {

	res := &TreeNode{}
	res.tag = Tree_node_tag_skip
	res.children = 0
	res.name = "empty"
	res.expr_or_state = -1

	empty_tree_node = res
}

// error node
func ErrorTreeNode(msg string) *TreeNode {

	res := &TreeNode{}
	res.tag = Tree_node_tag_erroneous
	res.children = 0
	res.name = "error:" + msg
	res.expr_or_state = -1

	return res
}
