package parser

import "husd.com/v0/ast_tree"

/**
 * Specifies the methods to access a mappings of syntax trees to end positions.
 * @author hushengdong
 */
type EndPosTable interface {
	/**
	 * This method will return the end position of a given tree, otherwise a
	 * Positions.NOPOS will be returned.
	 * @param jcTree JCTree
	 * @return position of the source tree or Positions.NOPOS for non-existent mapping
	 */
	GetEndPos(jcTree *ast_tree.JCTree) int
	/**
	 * Store ending position for a tree, the value of which is the greater of
	 * last error position and the given ending position.
	 * @param JCTree The tree.
	 * @param endPos The ending position to associate with the tree.
	 */
	SetEnd(jcTree *ast_tree.JCTree, endPos int)
	/**
	 * Give an old tree and a new tree, the old tree will be replaced with
	 * the new tree, the position of the new tree will be that of the old
	 * tree.
	 * @param oldTree a JCTree to be replaced
	 * @param newTree a JCTree to be replaced with
	 * @return position of the old tree or Positions.NOPOS for non-existent mapping
	 */
	ReplaceTree(oldTree *ast_tree.JCTree, newTree *ast_tree.JCTree) int
}
